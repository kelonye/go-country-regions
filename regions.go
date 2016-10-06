package regions

import (
  "fmt"
  "encoding/json"
)

var regions map[string] string

func init() {

  if data, err := Asset("regions.json"); err != nil {
    fmt.Printf("error reading file: %v\n", err)
  } else {
    if err := json.Unmarshal(data, &regions); err != nil {
      fmt.Printf("error parsing json: %v\n", err)
    }
  }

}

func Region(country string) (string, bool){

  region := regions[country]

  return region, region != ""

}
