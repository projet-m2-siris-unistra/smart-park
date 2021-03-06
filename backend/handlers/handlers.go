package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"reflect"

	"github.com/nats-io/nats.go"
)

var handlers = map[string]interface{}{
	"ping":                    ping,                  // ping database server
	"devices.get":             getDevice,             // get a device by its id
	"devices.get.notassigned": getNotAssignedDevices, // get all not assigned devices
	"devices.get.free":        getFreeDevices,        // get all free devices
	"tenants.get":             getTenant,             // get a tenant by its id
	"zones.get":               getZone,               // get a zone by its id
	"places.get":              getPlace,              // get a places by its id
	"users.get":               getUser,               // get a user by his/her id
	"devices.list":            getDevices,            // get all devices
	"tenants.list":            getTenants,            // get all tenants
	"zones.list":              getZones,              // get all zones by the tenant's id
	"places.list":             getPlaces,             // get all places by the zones' id
	"users.list":              getUsers,              // get all users
	"devices.update":          updateDevice,          // update devices' field
	"tenants.update":          updateTenants,         // update tenants' field
	"zones.update":            updateZone,            // update zones' field
	"places.update":           updatePlace,           // update places' field
	"users.update":            updateUser,            // update users' field
	"devices.new":             newDevice,             // create new device
	"places.new":              newPlace,              // create new place
	"zones.new":               newZone,               // create new zone
	"places.delete":           deletePlace,           // remove the place
	"zones.delete":            deleteZone,            // remove the zone and all places
	"devices.delete":          deleteDevice,          // remove the device
	"faker.new":               createFakeData,        // create fake data into the database (all tables)
}

// wrapHandler wraps a handler to do error handling and request/reply marshaling/unmarshaling
//
// It also injects a context which should be used throughout the handler.
// This is done by using `reflect` on the handler.
//
// The handler can look like this: func(ctx context.Context, req MyRequestType) (MyReplyType, error)
//
// TODO: better error handling
func wrapHandler(fn interface{}) func(*nats.Msg) {
	// First, do some reflection on the handler.
	// We especially need the type of the second argument and the type of the
	// first return value.
	handler := reflect.ValueOf(fn)
	fnType := reflect.TypeOf(fn)

	// Check the number of inputs variables
	// If there is only one, this topic does not have any argyments
	// If there are two, the message will be parsed with the type of the first output
	var hasRequest bool
	var requestType reflect.Type
	switch fnType.NumIn() {
	case 1:
		hasRequest = false
	case 2:
		hasRequest = true
		requestType = fnType.In(1)
	default:
		log.Fatal(errors.New("invalid number of inputs in handler"))
	}

	// Check the number of output variables
	// If there is only one, this topic does not reply anything
	// If there are two, this topic will reply with the type of the first output
	var hasReply bool
	var replyType reflect.Type
	switch fnType.NumOut() {
	case 1:
		hasReply = false
	case 2:
		hasReply = true
		replyType = fnType.Out(0)
	default:
		log.Fatal(errors.New("invalid number of outputs in handler"))
	}

	// Here's the "real" handler
	return func(m *nats.Msg) {
		var err error

		// Create the context for this handler
		// We might want to inherit it from somewhere later
		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()

		ctxValue := reflect.ValueOf(ctx)

		// Prepare the arguments for the call to the handler
		callArgs := []reflect.Value{ctxValue}

		if hasRequest {
			// Create an empty request type from the handler signature
			request := reflect.New(requestType).Interface()

			// Unmarshal the JSON message
			err = json.Unmarshal(m.Data, request)
			if err != nil {
				log.Println(err)
				return
			}
			requestValue := reflect.ValueOf(request)

			callArgs = append(callArgs, reflect.Indirect(requestValue))
		}

		// Call the handler.
		// Calling a reflected function is a bit tricky and involves using
		// reflect.Value objects.
		out := handler.Call(callArgs)

		// Extract the error and check it.
		// It is the last thing returned by the handler.
		errValue := out[len(out)-1]
		if !errValue.IsNil() {
			err = errValue.Interface().(error)
			log.Println(err)
			return
		}

		// If this handler does reply, marshal the reply and send it
		if hasReply {
			reply := out[0].Convert(replyType).Interface()

			payload, err := json.Marshal(reply)

			if err != nil {
				log.Println(err)
				return
			}

			err = m.Respond(payload)

			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

// Register the handlers
func Register(conn *nats.Conn) {
	log.Println("handlers: register")
	for name, fn := range handlers {
		log.Printf("handlers: subscribing to %s", name)
		conn.QueueSubscribe(name, "backend", wrapHandler(fn))
	}
}
