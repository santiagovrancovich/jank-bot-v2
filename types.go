package main

type Config struct {
	Dni       string `json:"dni"`
	Clave     string `json:"clave"`
	Comedores []struct {
		Nombre     string   `json:"nombre"`
		HoraInicio string   `json:"horaInicio"`
		Dias       []string `json:"dias"`
		Comida     string   `json:"comida"`
		ParaLlevar bool     `json:"paraLlevar"`
	} `json:"comedores"`
	SleepTime     int    `json:"sleepTime"`
	Concurrent    bool   `json:"concurrent"`
	MaxRandomTime int    `json:"maxRandomTime"`
	Token         string `json:"token"`
	Channel       int    `json:"channel"`
}
