module github.com/projet-m2-siris-unistra/smart-park/api

go 1.13

require (
	github.com/gorilla/mux v1.7.3
	github.com/nats-io/nats.go v1.9.1
	github.com/projet-m2-siris-unistra/smart-park/backend v0.0.0-20191224143719-86d60215529d
	gopkg.in/guregu/null.v3 v3.4.0
)

replace github.com/projet-m2-siris-unistra/smart-park/backend => ../backend
