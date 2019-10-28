package main

import (
	"fmt"
	"log"
	"os"
)

func publish(subj string, data string) error {
	fmt.Printf("Publishing to: %s\n", subj)
	err := nc.Publish(subj, []byte(data))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return nc.Flush()
}
