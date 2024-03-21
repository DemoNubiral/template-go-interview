package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/asim/go-micro/v3/web"
	"nubiral.com/template-go-micro/internal/source"
)

func main() {

	service := web.NewService(
		web.Name("template.service"),
		web.Address(":8080"),
	)

	service.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		cmd := source.GetHelloCMD{
			Parameter1: r.URL.Query().Get("param1"),
			Parameter2: r.URL.Query().Get("param2"),
		}

		response, _ := source.Execute(cmd)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Iniciar el servicio
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
