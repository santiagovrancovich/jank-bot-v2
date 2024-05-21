package main

import (
	"net/http"
	"slices"
	"time"
)

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

func checkTurnos(client *http.Client, servicio Servicio, dias []string) []Turno {
	var Turnos []Turno
	date := time.Now()
	nextWeek := date.Add(time.Hour * 24 * 7)

	if date.Month() == nextWeek.Month() {
		Turnos = buscarTurnos(client, servicio, date)
	} else {
		Turnos = slices.Concat(
			buscarTurnos(client, servicio, date),
			buscarTurnos(client, servicio, nextWeek),
		)
	}

	return filtrarDias(Turnos, dias)
}
