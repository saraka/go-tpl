package main

import (
	"log"
	"{{.ModuleName}}/internal/app/{{.ProjectName}}"
)

func main() {
	if err := {{.ProjectName}}.Run(); err != nil {
		log.Fatal(err)
	}
}
