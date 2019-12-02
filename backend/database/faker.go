package database

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

// Faker : insert fake data into the database
func Faker(ctx context.Context, tenants int, zones int, devices int, places int, users int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	gofakeit.Seed(0)

	if tenants > 0 {
		for n := 0; n <= tenants; n++ {
			nameTenant := gofakeit.City()
			GeoTenant := NewGeos()
			_, err := pool.ExecContext(ctx,
				`INSERT INTO tenants
				(
					name, 
					geo
				) VALUES
				(
					$1,
					$2
				)`, nameTenant, GeoTenant)

			if err != nil {
				return errors.New("error new device")
			}
		}
	}
	if zones > 0 {
		for n := 0; n <= zones; n++ {
			tenantIDZone := RandomTenantRow(ctx)
			nameZone := gofakeit.StreetName()
			typeZone := TypeZoneRandom()
			colorZone, err := RandomHex(6)
			colorZone = strings.ToUpper(colorZone)
			colorZone = colorZone[0:6]
			geoZone := NewGeos()
			_, err = pool.ExecContext(ctx,
				`INSERT INTO zones
				(
					tenant_id, 
					name,
					type,
					color,
					geo
				) VALUES
				(
					$1,
					$2,
					$3,
					$4,
					$5
				)`, tenantIDZone, nameZone, typeZone, colorZone, geoZone)

			if err != nil {
				return errors.New("error new zone")
			}
		}
	}

	if devices > 0 {
		for n := 0; n <= devices; n++ {
			batteryDevice := rand.Intn(100)
			stateDevice := StateDeviceRandom()
			_, err := pool.ExecContext(ctx,
				`INSERT INTO 
				devices (
					battery, 
					state
				) VALUES (
					$1,
					$2
				)`, batteryDevice, stateDevice)

			if err != nil {
				return errors.New("error new device")
			}
		}
	}

	if places > 0 {
		for n := 0; n <= places; n++ {
			deviceIDPlace := RandomDeviceRow(ctx)
			zoneIDPlace := RandomZoneRow(ctx)
			typePlace := ""
			geoPlace := NewGeo()
			_, err := pool.ExecContext(ctx,
				`INSERT INTO places
				(
					zone_id, 
					type,
					geo,
					device_id
				) VALUES
				(
					$1,
					$2,
					$3,
					$4
				)`, zoneIDPlace, typePlace, geoPlace, deviceIDPlace)

			if err != nil {
				return errors.New("error new place")
			}
		}
	}

	if users > 0 {
		for n := 0; n <= users; n++ {
			tenantIDUser := RandomTenantRow(ctx)
			nameUser := strings.ToLower(gofakeit.Username())
			passwordUser := gofakeit.Password(true, true, true, false, false, 1)
			emailUser := gofakeit.Email()
			_, err := pool.ExecContext(ctx,
				`INSERT INTO users
				(
					tenant_id,
					username,
					password,
					email
				) VALUES
				(
					$1,
					$2,
					$3,
					$4
				)`, tenantIDUser, nameUser, passwordUser, emailUser)

			if err != nil {
				return errors.New("error new user")
			}
		}
	}

	return nil
}

// NewGeos : create a string which has a list of coordinates
func NewGeos() string {
	var result string
	result = "["
	for n := 0; n <= rand.Intn(10); n++ {
		tmpLongitude := fmt.Sprintf("%f", gofakeit.Longitude())
		tmpLatitude := fmt.Sprintf("%f", gofakeit.Latitude())
		result = result + "[" + tmpLongitude + "," + tmpLatitude + "],"
	}
	result = result + "]"
	return result
}

// NewGeo : create a string which has coordinates
func NewGeo() string {
	var result string
	tmpLongitude := fmt.Sprintf("%f", gofakeit.Longitude())
	tmpLatitude := fmt.Sprintf("%f", gofakeit.Latitude())
	result = result + "[" + tmpLongitude + "," + tmpLatitude + "],"
	return result
}

// TypeZoneRandom : return a type zone
func TypeZoneRandom() string {
	n := rand.Intn(2)
	if n == 0 {
		return "paid"
	} else if n == 1 {
		return "blue"
	} else {
		return "free"
	}
}

// StateDeviceRandom : return a type zone
func StateDeviceRandom() string {
	n := rand.Intn(1)
	if n == 0 {
		return "occupied"
	}

	return "free"
}

// RandomTenantRow : return a random row from the tenant table
func RandomTenantRow(ctx context.Context) int {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var id int

	err := pool.QueryRowContext(ctx, `
		SELECT tenant_id
		FROM tenants ORDER BY RANDOM()
		LIMIT 1`).Scan(&id)

	if err != nil {
		return -1
	}

	return id
}

// RandomDeviceRow : return a random row from the device table
func RandomDeviceRow(ctx context.Context) int {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var id int

	err := pool.QueryRowContext(ctx, `
		SELECT device_id
		FROM devices ORDER BY RANDOM()
		LIMIT 1`).Scan(&id)

	if err != nil {
		return -1
	}

	return id
}

// RandomZoneRow : return a random row from the zone table
func RandomZoneRow(ctx context.Context) int {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var id int

	err := pool.QueryRowContext(ctx, `
		SELECT zone_id
		FROM zones ORDER BY RANDOM()
		LIMIT 1`).Scan(&id)

	if err != nil {
		return -1
	}

	return id
}

// RandomHex : generate random hex string
func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
