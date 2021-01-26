# DO Metadata
[![godocs.io](https://godocs.io/github.com/codegaudi/do-metadata?status.svg)](https://godocs.io/github.com/codegaudi/do-metadata)

This package implements a simple way to retrieve droplet metadata from the relevant digitalocean endpoint.

## Example
```golang
package main

import (
	do_metadata "github.com/codegaudi/do-metadata"
	"log"
	"time"
)

func main() {

    meta, err := do_metadata.RetrieveMetadata(time.Second * 5)
    if err != nil {
        log.Fatal("could not retrieve metadata from the digitalocean metadata api: ", err.Error())
    }
    
    // Print the droplet id
    log.Print(meta.DropletID)
}
```

## License
MIT