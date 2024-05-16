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
	"strings"
	"time"
)

var conf Config

func buscarTurnos(client *http.Client, s Servicio) []Turno {
	var servicioHoy ServicioDia

	servicioHoy.Servicio = s
	servicioHoy.Fecha = fmt.Sprintf("%s 00:00:00", time.Now().Format(time.DateOnly))

	horaInicio := servicioHoy.Servicio.HoraInicio.(map[string]interface{})
	horaFin := servicioHoy.Servicio.HoraFin.(map[string]interface{})

	servicioHoy.Servicio.HoraInicio = horaInicio["horaCorta"]
	servicioHoy.Servicio.HoraFin = horaFin["horaCorta"]

	strBytes, err := json.Marshal(servicioHoy)

	if err != nil {
		fmt.Println(err)
	}

	str := url.QueryEscape(string(strBytes))
	body := strings.NewReader(fmt.Sprintf("json=%s", str))

	resp, err := client.Post("https://comedores.unr.edu.ar/comedor-reserva/buscar-turnos-reservas", "application/x-www-form-urlencoded", body)
	turnos, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var t TurnoRequest
	json.Unmarshal(turnos, &t)

	var TurnosArray []Turno

	for _, i := range t.Turnos {
		TurnosArray = append(TurnosArray, i)
	}

	return TurnosArray
}

func filtrarServicios(comedores []Comedor) []Servicio {
	var Servicios []Servicio

	for _, comedorConf := range conf.Comedores {
		for _, comedor := range comedores {
			if comedor.Nombre == comedorConf.Nombre {
				for _, servicio := range comedor.Servicios {
					if servicio.HoraInicio.(map[string]interface{})["horaCorta"] == comedorConf.HoraInicio &&
						servicio.Tipo.Nombre == comedorConf.Comida &&
						servicio.ParaLlevar == comedorConf.ParaLlevar {
						Servicios = append(Servicios, servicio)
					}
				}
			}
		}
	}

	return Servicios
}

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

	Turnos := buscarTurnos(client, filtrarServicios(comedoresArray)[0])
	fmt.Println(Turnos[0])
}
