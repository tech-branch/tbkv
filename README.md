## TBKV key-val store

You're looking at a simple zero-dependency in-memory key-value store for Golang. 

All operations (Get/Set/Delete) are executed in a single, buffered loop, so the store provides some thread safety.

### Usage


```Golang
kvs := NewStore()

kvs.Set("key", "value")

val, err := kvs.Get("key")
if err != nil {
    return err
}

// val == "value"

kvs.Delete("key")

val, err := kvs.Get("key")
if err != nil {
    // err == ErrNotFound
    return err
}
```
