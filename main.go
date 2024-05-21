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
	"slices"
	"strings"
	"time"
)

var conf Config

func pedirTurno(client *http.Client, turnos []Turno) {
	for _, turno := range turnos {
		if turno.Reserva != nil && time.Now().Format(time.DateTime) >= turno.Fecha.FechaMysql {
			strId := strings.NewReader(fmt.Sprintf("turno=%d", turno.ID))
			resp, err := client.Post("https://comedores.unr.edu.ar/comedor-reserva/buscar-turnos-reservas", "application/x-www-form-urlencoded", strId)

			if err != nil {
				fmt.Println("Invalid ID")
			}

			fmt.Printf("REQUEST ID:%d %s Status: %d\n", turno.ID, turno.Fecha.Fecha, resp.StatusCode)
			defer resp.Body.Close()
		}
	}
}

func buscarTurnos(client *http.Client, s Servicio, f time.Time) []Turno {
	var servicioHoy ServicioDia

	servicioHoy.Servicio = s
	servicioHoy.Fecha = fmt.Sprintf("%s 00:00:00", f.Format(time.DateOnly))

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
	TurnosArray = append(TurnosArray, t.Turnos...)

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

func filtrarDias(Turnos []Turno, Days []string) []Turno {
	var t []Turno

	for _, turno := range Turnos {
		if slices.Contains(Days, turno.Fecha.DiaNombre[0:2]) {
			t = append(t, turno)
		}
	}

	return t
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

	arrServicios := filtrarServicios(comedoresArray)
	date := time.Now()
	nextWeek := date.Add(time.Hour * 24 * 7)

	for i, servicio := range arrServicios {
		var Turnos []Turno

		if date.Month() == nextWeek.Month() {
			Turnos = buscarTurnos(client, servicio, date)
		} else {
			Turnos = slices.Concat(
				buscarTurnos(client, servicio, date),
				buscarTurnos(client, servicio, nextWeek),
			)
		}

		Turnos = filtrarDias(Turnos, conf.Comedores[i].Dias)

		pedirTurno(client, Turnos)
	}
}
