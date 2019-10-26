package main

import (
	"log"
	"sync"

	"github.com/nats-io/go-nats"
)

func subscribe(subj string) error {
	// Use a WaitGroup to wait for a message to arrive - 1 message
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe to the updates
	if _, err := nc.QueueSubscribe("*", "queue", func(msg *nats.Msg) {
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()

	// [end subscribe_async]
	nc.Close()

	return nil
}
