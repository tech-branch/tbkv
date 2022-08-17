package main

import (
	"log"

	"github.com/tech-branch/tbkv"
)

// testing a simple Set/Get/Delete/Get flow
func main() {
	kvs := tbkv.NewDefaultStore()

	const (
		key   = "key1"
		value = "value1"
	)

	// ---
	// Set
	// ---

	log.Printf("Setting key ' %s ' to ' %s ' ", key, value)
	kvs.Set(key, value)

	// ---
	// Get
	// ---

	log.Printf("Fetching key ' %s ' ", key)
	val, err := kvs.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got value: %s", val)

	// ------
	// Delete
	// ------

	log.Printf("Deleting key ' %s '", key)
	kvs.Delete(key)

	// ---
	// Get
	// ---
	// Expect ErrNotFound because key does not exist

	log.Printf("Fetching key ' %s '", key)
	_, err = kvs.Get(key)
	if err == tbkv.ErrNotFound {
		log.Print("Encountered expected ErrNotFound error")
	} else {
		log.Fatal("Expected ErrNotFound error, got ", err)
	}
}
