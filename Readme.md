### Install

    $ go get github.com/kelonye/go-country-regions

### Use

```go

package main

import (
  "fmt"
	"github.com/kelonye/go-country-regions"
)

func main(){

	region, _ := regions.Region("US")

	fmt.Printf("%s\n", region)

}

```
