package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zeebo/bencode"
	"gopkg.in/yaml.v2"
)

func main() {
	var m map[string]interface{}
	if err := bencode.NewDecoder(os.Stdin).Decode(&m); err != nil {
		log.Fatalf("bencode error: %v", err)
	}

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("yaml error: %v", err)

	}
	fmt.Print(string(d))
}
