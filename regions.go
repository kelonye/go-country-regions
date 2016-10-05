package regions

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
)

var regions map[string] string

func init() {

  if file, err := ioutil.ReadFile("./regions.json"); err != nil {
    fmt.Errorf("error reading file: %v\n", err)
  } else {
    if err := json.Unmarshal(file, &regions); err != nil {
      fmt.Errorf("error parsing json: %v\n", err)
    }
  }

}

func Region(country string) (string, bool){

  region := regions[country]

  return region, region != ""

}
