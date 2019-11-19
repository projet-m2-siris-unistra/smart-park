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
	"ping":         ping,
	"devices.get":  getDevice,
	"tenants.get":  getTenant,
	"zones.get":    getZone,
	"places.get":   getPlace,
	"users.get":    getUser,
	"devices.list": getDevices,
	//"tenants.get": getTenants,
	//"zones.get": getZones,
	//"places.get": getPlaces,
	//"users.get": getUsers,
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
		conn.Subscribe(name, wrapHandler(fn))
	}
}
