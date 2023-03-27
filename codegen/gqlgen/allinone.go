// go:build ignore
package main

import (
	"github.com/99designs/gqlgen/codegen/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.LoadConfig("./codegen/gqlgen/gqlgen.yaml")
	if err != nil {
		log.Fatal("failed to load config", err.Error())
	}
	// clear file
	err = os.WriteFile("./test/testdata/allinone.graphql", []byte(""), os.ModePerm)
	if err != nil {
		log.Fatal("failed to write file", err.Error())
	}
	f, err := os.OpenFile("./test/testdata/allinone.graphql", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("failed to write file", err.Error())
	}
	defer f.Close()
	for _, source := range cfg.Sources {
		// write to file
		_, err = f.WriteString(source.Input)
		if err != nil {
			log.Fatal("failed to write file", err.Error())
		}
	}
}
