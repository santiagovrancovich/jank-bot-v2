package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"time"
)

var conf Config

func getComedores(client *http.Client) []Comedor {
	resp, err := client.Get("https://comedores.unr.edu.ar/comedor-reserva/reservar")
	reg, _ := regexp.Compile("var jsonReservar[\\s\\S]*?\\};")

	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	jsonBytes := reg.Find(b)
	var jsonResp ReservarJson
	json.Unmarshal(jsonBytes[19:len(jsonBytes)-1], &jsonResp)

	var comedoresArray []Comedor
	for _, confComedor := range conf.Comedores {
		for _, jsonComedor := range jsonResp.Comedores {
			if confComedor.Nombre == jsonComedor.Nombre {
				comedoresArray = append(comedoresArray, jsonComedor)
			}
		}
	}

	return comedoresArray
}

func main() {
	file, err := os.ReadFile("config.json")

	if err != nil {
		log.Fatalln("Missing or invalid config file")
	}

	err = json.Unmarshal(file, &conf)

	jar, err := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	// Get site cookie and auth
	client.Get("https://comedores.unr.edu.ar/")
	client.PostForm("https://comedores.unr.edu.ar/", url.Values{
		"form-login[dni]":      {conf.Dni},
		"form-login[clave]":    {conf.Clave},
		"botones[botonEnviar]": {},
	})

	comedoresArray := getComedores(client)

	for len(conf.Comedores) != len(comedoresArray) {
		time.Sleep(time.Millisecond * time.Duration(conf.SleepTime))
		comedoresArray = getComedores(client)
	}

	fmt.Println(comedoresArray)
}
