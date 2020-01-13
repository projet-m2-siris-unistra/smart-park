package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/nats-io/nats.go"
	"golang.org/x/oauth2"
)

func loadTLS(prefix string) (*tls.Config, error) {
	certFile, ok := os.LookupEnv(prefix + "_CERT")
	if !ok {
		return nil, nil
	}
	keyFile, ok := os.LookupEnv(prefix + "_KEY")
	if !ok {
		return nil, nil
	}
	caFile, ok := os.LookupEnv(prefix + "_CA")
	if !ok {
		return nil, nil
	}

	// Load client certificate
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// Load CA
	pool := x509.NewCertPool()
	rootPEM, err := ioutil.ReadFile(caFile)
	if err != nil || rootPEM == nil {
		return nil, fmt.Errorf("error loading or parsing rootCA file: %v", err)
	}
	ok = pool.AppendCertsFromPEM(rootPEM)
	if !ok {
		return nil, fmt.Errorf("failed to parse root certificate from %q", caFile)
	}

	return &tls.Config{
		MinVersion:   tls.VersionTLS12,
		RootCAs:      pool,
		Certificates: []tls.Certificate{cert},
	}, nil
}

var conn *nats.Conn
var jsonConn *nats.EncodedConn

func busInit(url string, tlsConfig *tls.Config) error {
	log.Println("bus: init")
	var err error

	if tlsConfig != nil {
		conn, err = nats.Connect(url, nats.Secure(tlsConfig))
	} else {
		conn, err = nats.Connect(url)
	}
	if err != nil {
		return err
	}

	jsonConn, err = nats.NewEncodedConn(conn, nats.JSON_ENCODER)

	return err
}

type providerConfig struct {
	TenantID     string `json:"tenant_id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Issuer       string `json:"issuer"`
}

func loadProviderConfigs(path string) (*map[string]providerConfig, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var result map[string]providerConfig
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type provider struct {
	*oauth2.Config
	*oidc.Provider
	*oidc.IDTokenVerifier
	tenantID int
	iss      string
}

var providers map[string]*provider

type authorizationRequest struct {
	Config      string `json:"config"`
	RedirectURI string `json:"redirect_uri"`
}

type authorizationResponse struct {
	URL string `json:"url"`
}

type validateRequest struct {
	Token string `json:"token"`
}

type partialToken struct {
	ISS string `json:"iss"`
}

type user struct {
	Sub     string `json:"sub"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type validateResponse struct {
	Valid    bool  `json:"valid"`
	TenantID *int  `json:"tenant_id,omitempty"`
	User     *user `json:"user,omitempty"`
}

type exchangeRequest struct {
	Config      string `json:"config"`
	Code        string `code:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type exchangeResponse struct {
	Token string `json:"token"`
}

func authorizationHandler(subject, reply string, req *authorizationRequest) {
	if p, ok := providers[req.Config]; ok {
		url := p.Config.AuthCodeURL("", oauth2.SetAuthURLParam("redirect_uri", req.RedirectURI))
		resp := authorizationResponse{URL: url}
		err := jsonConn.Publish(reply, &resp)
		if err != nil {
			log.Printf("error while sending response: %v", err)
			return
		}
	}
}

func exchangeHandler(subject, reply string, req *exchangeRequest) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	if p, ok := providers[req.Config]; ok {
		token, err := p.Config.Exchange(ctx, req.Code, oauth2.SetAuthURLParam("redirect_uri", req.RedirectURI))
		if err != nil {
			log.Printf("error while exchanging code: %v", err)
			return
		}

		resp := exchangeResponse{Token: token.Extra("id_token").(string)}
		err = jsonConn.Publish(reply, resp)
		if err != nil {
			log.Printf("error while sending response: %v", err)
			return
		}
	}
}

func issFromJWT(token string) (*string, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("oidc: malformed jwt, expected 3 parts got %d", len(parts))
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("oidc: malformed jwt payload: %v", err)
	}

	claims := partialToken{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, fmt.Errorf("oidc: could not parse token: %v", err)
	}

	return &claims.ISS, nil
}

func findConfig(iss string) *provider {
	for _, p := range providers {
		if p.iss == iss {
			return p
		}
	}
	return nil
}

func validateHandler(subject, reply string, req *validateRequest) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	var p *provider
	resp := validateResponse{Valid: false}
	claims := user{}

	iss, err := issFromJWT(req.Token)
	if err != nil {
		log.Print(err)
		goto end
	}

	p = findConfig(*iss)
	if p == nil {
		log.Printf("failed to find provider for iss %s", *iss)
		goto end
	}

	if idToken, err := p.IDTokenVerifier.Verify(ctx, req.Token); err == nil {
		if err := idToken.Claims(&claims); err != nil {
			log.Printf("failed to extract claims from token: %v", err)
			goto end
		}
	} else {
		goto end
	}

	resp.Valid = true
	resp.TenantID = &p.tenantID
	resp.User = &claims

end:
	jsonConn.Publish(reply, &resp)
}

func loadProvider(ctx context.Context, tenantID int, iss, clientID, clientSecret string) (*provider, error) {
	oidcProvider, err := oidc.NewProvider(ctx, iss)
	if err != nil {
		return nil, fmt.Errorf("unable to init oidc provider: %v", err)
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     oidcProvider.Endpoint(),
		Scopes:       []string{"openid", "profile"},
	}

	verifier := oidcProvider.Verifier(&oidc.Config{ClientID: clientID})

	return &provider{
		Provider:        oidcProvider,
		Config:          &config,
		IDTokenVerifier: verifier,
		tenantID:        tenantID,
		iss:             iss,
	}, nil
}

func main() {
	ctx := context.Background()
	natsURL, ok := os.LookupEnv("NATS_URL")
	if !ok {
		natsURL = nats.DefaultURL
	}

	tlsConfig, err := loadTLS("NATS")
	if err != nil {
		log.Fatalf("unable to load certificates: %v", err)
	}

	err = busInit(natsURL, tlsConfig)
	if err != nil {
		log.Fatalf("unable to connect to bus: %v", err)
	}
	defer conn.Close()

	providersConfigPath, ok := os.LookupEnv("PROVIDERS_CONFIG")
	if !ok {
		log.Fatal("no PROVIDERS_CONFIG")
	}

	configs, err := loadProviderConfigs(providersConfigPath)
	if err != nil {
		log.Fatalf("unable to load providers: %v", err)
	}

	providers = map[string]*provider{}
	for key, value := range *configs {
		tenantID, err := strconv.Atoi(value.TenantID)
		if err != nil {
			log.Fatalf("invalid tenant_id '%s': %v", value.TenantID, err)
		}

		p, err := loadProvider(ctx, tenantID, value.Issuer, value.ClientID, value.ClientSecret)
		if err != nil {
			log.Fatalf("failed to setup provider %s: %v", key, err)
		}
		providers[key] = p
	}

	jsonConn.Subscribe("auth.authorization", authorizationHandler)
	jsonConn.Subscribe("auth.exchange", exchangeHandler)
	jsonConn.Subscribe("auth.validate", validateHandler)

	// wait signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Receive from c
	log.Println("main: exiting")
}
