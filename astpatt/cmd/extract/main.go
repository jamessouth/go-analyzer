package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jamessouth/go-analyzer/astpatt"
	"github.com/jamessouth/go-analyzer/astrav"
)

var (
	path = flag.String("path", ".", "path to example solution folder")
)

func main() {
	flag.Parse()

	folder := astrav.NewFolder(http.Dir(*path), "")
	packages, err := folder.ParseFolder()
	if err != nil {
		log.Fatal(err)
	}

	for _, pkg := range packages {
		pattern := astpatt.ExtractPattern(pkg)
		data, err := json.MarshalIndent(pattern, "", "  ")
		fmt.Println(string(data))

		if err != nil {
			log.Println(err)
		}
	}

}
