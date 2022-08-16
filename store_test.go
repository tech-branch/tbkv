package tbkv

import (
	"fmt"
	"testing"
)

func TestKVStore_Set(t *testing.T) {
	// testing the usual Set/Get flow
	kvs := NewDefaultStore()
	kvs.Set("key", "value")
	val, err := kvs.Get("key")
	if err != nil {
		t.Error(err)
	}
	if val != "value" {
		t.Errorf("Expected value to be 'value', got '%s'", val)
	}
}

func TestKVStore_SetLargeBuffer(t *testing.T) {
	// testing the same Set/Get flow but with a bigger buffer
	kvs := NewStore(255)
	kvs.Set("key", "value")
	val, err := kvs.Get("key")
	if err != nil {
		t.Error(err)
	}
	if val != "value" {
		t.Errorf("Expected value to be 'value', got '%s'", val)
	}
}

func TestKVStore_DoubleSet(t *testing.T) {
	// I want to overwrite entries without an issue
	kvs := NewDefaultStore()
	kvs.Set("key", "value1")
	kvs.Set("key", "value2")
}

func TestKVStore_SetMultiple(t *testing.T) {
	// lets try with multiple consecutive sets
	kvs := NewDefaultStore()
	for i := 0; i < 20; i++ {
		kvs.Set(fmt.Sprint(i), fmt.Sprint(i*i))
	}
	val, err := kvs.Get("19")
	if err != nil {
		t.Error(err)
	}
	if val != "361" {
		t.Errorf("Expected value to be '361', got '%s'", val)
	}
}

func TestKVStore_Delete(t *testing.T) {
	// I want to delete entries without an issue
	// I don't want to handle an error if entry does not exist, assume Ok
	kvs := NewDefaultStore()
	for i := 0; i < 20; i++ {
		kvs.Set(fmt.Sprint(i), fmt.Sprint(i*i))
	}
	kvs.Delete("19")
	// this is a no-op because the key doesn't exist
	kvs.Delete("19")

	_, err := kvs.Get("19")
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got '%s'", err)
	}

}
