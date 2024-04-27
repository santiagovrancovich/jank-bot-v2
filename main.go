package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func main() {
	file, err := os.ReadFile("config.json")

	if err != nil {
		log.Fatalln("Missing or invalid config file")
	}

	var conf Config
	err = json.Unmarshal(file, &conf)

	jar, err := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	// Get site cookie and auth
	resp, err := client.Get("https://comedores.unr.edu.ar/")
	client.PostForm("https://comedores.unr.edu.ar/", url.Values{
		"form-login[dni]":      {conf.Dni},
		"form-login[clave]":    {conf.Clave},
		"botones[botonEnviar]": {},
	})

	fmt.Println(resp.Cookies())
}
