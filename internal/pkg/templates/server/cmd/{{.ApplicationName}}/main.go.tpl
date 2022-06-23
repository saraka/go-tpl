package main

import (
	"log"
	"{{.ModuleName}}/internal/app/{{.ApplicationName}}"
)

func main() {
	if err := {{.ApplicationName}}.Run(); err != nil {
		log.Fatal(err)
	}
}
