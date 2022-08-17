![tests](https://github.com/tech-branch/tbkv/actions/workflows/main.yml/badge.svg)
![continuous semgrep](https://github.com/tech-branch/tbkv/actions/workflows/semgrep.yml/badge.svg)

## TBKV key-val store

You're looking at a simple zero-dependency in-memory key-value store for Golang. 

All operations (Get/Set/Delete) are executed in a single, buffered loop, so the store provides some thread safety.

The package is continuously scanned by `semgrep` and various static code critics to maintain healthy security posture.

The main `tbkv` package has 100% test coverage and my aim is to maintain that. 
The `examples` are covered to the extent of checking if the scenario runs successfully. 

### Usage


```Golang

import (
	"github.com/tech-branch/tbkv"
)

func main() {
    kvs := NewDefaultStore()

    kvs.Set("key", "value")

    val, err := kvs.Get("key")
    if err != nil {
        return err
    }

    // val == "value"

    kvs.Delete("key")

    val, err := kvs.Get("key")
    if err != nil {
        // err == tbkv.ErrNotFound
        return err
    }
}
```

See a more complete example in `./examples/main.go`
