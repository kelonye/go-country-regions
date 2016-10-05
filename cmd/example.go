package main

import (
  "fmt"
	"github.com/kelonye/go-country-regions"
)

func main(){

	region, _ := regions.Region("US")

	fmt.Printf("%s\n", region)

}
