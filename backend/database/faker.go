package database

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

// define global variables to delimiting the geo's tenant
var minlatitude float64
var maxlatitude float64
var minlongitude float64
var maxlongitude float64

// Faker : insert fake data into the database
func Faker(ctx context.Context, tenants int, zones int, devices int, places int, users int) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	gofakeit.Seed(0)

	if tenants > 0 {
		for n := 0; n <= tenants; n++ {
			nameTenant := gofakeit.City()
			GeoTenant, errgeo := NewGeoTenant()
			if errgeo != nil {
				return errors.New("error new geo's tenant, function Faker, faker.go")
			}
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
				return errors.New("error new tenant, function Faker, faker.go")
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
			geoZone, errgeo := NewGeoZone()
			if errgeo != nil {
				return errors.New("error new geo's zone, function Faker, faker.go")
			}
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
				return errors.New("error new zone, function Faker, faker.go")
			}
		}
	}

	if devices > 0 {
		for n := 0; n <= devices; n++ {
			batteryDevice := Random(0, 101)
			stateDevice := StateDeviceRandom()
			tenantID4Device := RandomTenantRow(ctx)
			deviceEUI := RandomEUIGenerator()
			_, err := pool.ExecContext(ctx,
				`INSERT INTO 
				devices (
					battery, 
					state,
					tenant_id,
					device_eui
				) VALUES (
					$1,
					$2,
					$3,
					$4
				)`, batteryDevice, stateDevice, tenantID4Device, deviceEUI)

			if err != nil {
				return errors.New("error new device, function Faker, faker.go")
			}
		}
	}

	if places > 0 {
		for n := 0; n <= places; n++ {
			deviceIDPlace := RandomDeviceRow(ctx)
			zoneIDPlace := RandomZoneRow(ctx)
			typePlace := ""
			geoPlace, errgeo := NewGeo()
			if errgeo != nil {
				log.Printf("query error: %v\n", errgeo)
				return errors.New("error new geo's place, function Faker, faker.go")
			}
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
				log.Printf("query error: %v\n", err)
				return errors.New("error new place, function Faker, faker.go")
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
				return errors.New("error new user, function Faker, faker.go")
			}
		}
	}

	return nil
}

// NewGeoTenant : create a string which has a list of coordinates for a tenant
func NewGeoTenant() (string, error) {
	var result string
	var err error
	var latitude float64
	var longitude float64
	result = "["

	// do a list of corrected approximative values of longitude
	minlongitude = gofakeit.Longitude()
	maxlongitude, err = gofakeit.LongitudeInRange(minlongitude, 180)
	if err != nil {
		return "", errors.New("error : wrong longitude, function NewGeoTenant, faker.go")
	}

	// do a list of corrected approximative values of latitude
	minlatitude = gofakeit.Latitude()
	maxlatitude, err = gofakeit.LatitudeInRange(minlatitude, 90)
	if err != nil {
		return "", errors.New("error : wrong latitude, function NewGeoTenant, faker.go")
	}

	for n := 0; n <= Random(2, 10); n++ {
		if n != 0 {
			result = result + ","
		}

		longitude, err = RandomFloat64(minlongitude, maxlongitude)
		if err != nil {
			return "", errors.New("error : wrong longitude, function NewGeoTenant, faker.go")
		}

		latitude, err = RandomFloat64(minlatitude, maxlatitude)
		if err != nil {
			return "", errors.New("error : wrong latitude, function NewGeoTenant, faker.go")
		}

		// format string and concat with result
		tmpLongitude := fmt.Sprintf("%f", longitude)
		tmpLatitude := fmt.Sprintf("%f", latitude)
		result = result + "[" + tmpLongitude + "," + tmpLatitude + "]"
	}
	result = result + "]"
	return result, nil
}

// NewGeo : create a string which has coordinates
func NewGeo() (string, error) {
	var result string
	var err error
	var latitude float64
	var longitude float64

	longitude, err = RandomFloat64(minlongitude, maxlongitude)
	if err != nil {
		return "", errors.New("error : wrong longitude, function NewGeo, faker.go")
	}

	latitude, err = RandomFloat64(minlatitude, maxlatitude)
	if err != nil {
		return "", errors.New("error : wrong latitude, function NewGeo, faker.go")
	}

	tmpLongitude := fmt.Sprintf("%f", longitude)
	tmpLatitude := fmt.Sprintf("%f", latitude)
	result = result + "[" + tmpLongitude + "," + tmpLatitude + "],"
	return result, nil
}

// NewGeoZone : create a string which has a list of coordinates for a zone and fit with the geo's tenant
func NewGeoZone() (string, error) {
	var result string
	var err error
	var latitude float64
	var longitude float64

	result = "["

	for n := 0; n <= Random(2, 10); n++ {
		longitude, err = RandomFloat64(minlongitude, maxlongitude)
		if err != nil {
			return "", errors.New("error : wrong longitude, function NewGeoZone, faker.go")
		}

		latitude, err = RandomFloat64(minlatitude, maxlatitude)
		if err != nil {
			return "", errors.New("error : wrong latitude, function NewGeoZone, faker.go")
		}

		// format string and concat with result
		tmpLongitude := fmt.Sprintf("%f", longitude)
		tmpLatitude := fmt.Sprintf("%f", latitude)
		result = result + "[" + tmpLongitude + "," + tmpLatitude + "],"
	}
	result = result + "]"
	return result, nil

}

// TypeZoneRandom : return a type zone
func TypeZoneRandom() string {
	n := Random(1, 4)
	if n == 1 {
		return "paid"
	} else if n == 2 {
		return "blue"
	} else {
		return "free"
	}
}

// StateDeviceRandom : return a type zone
func StateDeviceRandom() string {
	n := Random(1, 3)
	if n == 1 {
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

// Random : defines how to create the random number
func Random(min, max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(max-min) + min
}

// RandomFloat64 : defines how to create the random float64
func RandomFloat64(min, max float64) (float64, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return min + r1.Float64()*(max-min), nil
}

// RandomEUIGenerator : generates fake EUI
func RandomEUIGenerator() string {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		return ""
	}

	// Set the local bit
	buf[0] |= 2
	result := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
	return result[0:16]
}
