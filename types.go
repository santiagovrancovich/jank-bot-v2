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
type Comedor struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Foto      string `json:"foto"`
	Localidad struct {
		ID        int `json:"id"`
		Provincia struct {
			ID                    int           `json:"id"`
			Localidades           []interface{} `json:"localidades"`
			Capital               interface{}   `json:"capital"`
			Nombre                string        `json:"nombre"`
			Web                   string        `json:"web"`
			Bandera               string        `json:"bandera"`
			BanderaNombreOriginal string        `json:"banderaNombreOriginal"`
			Escudo                string        `json:"escudo"`
			EscudoNombreOriginal  string        `json:"escudoNombreOriginal"`
			Poblacion             int           `json:"poblacion"`
		} `json:"provincia"`
		CapitalPais      interface{} `json:"capitalPais"`
		CapitalProvincia interface{} `json:"capitalProvincia"`
		Nombre           string      `json:"nombre"`
		Web              string      `json:"web"`
		Bandera          string      `json:"bandera"`
		Escudo           string      `json:"escudo"`
		Poblacion        int         `json:"poblacion"`
		CodigoPostal     string      `json:"codigoPostal"`
		Latitud          float64     `json:"latitud"`
		Longitud         float64     `json:"longitud"`
	} `json:"localidad"`
	DireccionCalle  string `json:"direccionCalle"`
	DireccionNumero string `json:"direccionNumero"`
	FotoURL         string `json:"fotoUrl"`
	Servicios       []struct {
		ID      int `json:"id"`
		Comedor struct {
			ID     int    `json:"id"`
			Nombre string `json:"nombre"`
		} `json:"comedor"`
		Tipo struct {
			ID                     int    `json:"id"`
			Nombre                 string `json:"nombre"`
			AplicaBeneficiosBecado bool   `json:"aplicaBeneficiosBecado"`
			Orden                  int    `json:"orden"`
			Precios                []struct {
				CategoriaComensal struct {
					ID            int    `json:"id"`
					Nombre        string `json:"nombre"`
					NombreInterno string `json:"nombreInterno"`
				} `json:"categoriaComensal"`
				ID     int `json:"id"`
				Precio int `json:"precio"`
			} `json:"precios"`
		} `json:"tipo"`
		Dias []struct {
			Nombre     string `json:"nombre"`
			NumeroDia  int    `json:"numeroDia"`
			GetInicial string `json:"getInicial"`
		} `json:"dias"`
		Nombre             string `json:"nombre"`
		FechaVigenciaDesde struct {
			Hora       string `json:"hora"`
			HoraCorta  string `json:"horaCorta"`
			Fecha      string `json:"fecha"`
			Mysql      string `json:"mysql"`
			DiaNombre  string `json:"diaNombre"`
			FechaMysql string `json:"fechaMysql"`
			Timezone   struct {
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"timezone"`
			Timestamp  int    `json:"timestamp"`
			FechaCorta string `json:"fechaCorta"`
		} `json:"fechaVigenciaDesde"`
		FechaVigenciaHasta interface{} `json:"fechaVigenciaHasta"`
		HoraInicio         struct {
			Hora       string `json:"hora"`
			HoraCorta  string `json:"horaCorta"`
			Fecha      string `json:"fecha"`
			Mysql      string `json:"mysql"`
			DiaNombre  string `json:"diaNombre"`
			FechaMysql string `json:"fechaMysql"`
			Timezone   struct {
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"timezone"`
			Timestamp  int    `json:"timestamp"`
			FechaCorta string `json:"fechaCorta"`
		} `json:"horaInicio"`
		HoraFin struct {
			Hora       string `json:"hora"`
			HoraCorta  string `json:"horaCorta"`
			Fecha      string `json:"fecha"`
			Mysql      string `json:"mysql"`
			DiaNombre  string `json:"diaNombre"`
			FechaMysql string `json:"fechaMysql"`
			Timezone   struct {
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"timezone"`
			Timestamp  int    `json:"timestamp"`
			FechaCorta string `json:"fechaCorta"`
		} `json:"horaFin"`
		SinControlCupos               bool   `json:"sinControlCupos"`
		ParaLlevar                    bool   `json:"paraLlevar"`
		ComprobarTieneFechaHasta      bool   `json:"comprobarTieneFechaHasta"`
		ComprobarEsFutura             bool   `json:"comprobarEsFutura"`
		GetTiempoLimiteReservaMinutos int    `json:"getTiempoLimiteReservaMinutos"`
		Horario                       string `json:"horario"`
	} `json:"servicios"`
	ServiciosPorTipo []struct {
		Tipo struct {
			ID                     int    `json:"id"`
			Nombre                 string `json:"nombre"`
			AplicaBeneficiosBecado bool   `json:"aplicaBeneficiosBecado"`
			Orden                  int    `json:"orden"`
			Precios                []struct {
				CategoriaComensal struct {
					ID            int    `json:"id"`
					Nombre        string `json:"nombre"`
					NombreInterno string `json:"nombreInterno"`
				} `json:"categoriaComensal"`
				ID     int `json:"id"`
				Precio int `json:"precio"`
			} `json:"precios"`
		} `json:"tipo"`
		Servicios []struct {
			ID      int `json:"id"`
			Comedor struct {
				ID     int    `json:"id"`
				Nombre string `json:"nombre"`
			} `json:"comedor"`
			Tipo struct {
				ID                     int    `json:"id"`
				Nombre                 string `json:"nombre"`
				AplicaBeneficiosBecado bool   `json:"aplicaBeneficiosBecado"`
				Orden                  int    `json:"orden"`
				Precios                []struct {
					CategoriaComensal struct {
						ID            int    `json:"id"`
						Nombre        string `json:"nombre"`
						NombreInterno string `json:"nombreInterno"`
					} `json:"categoriaComensal"`
					ID     int `json:"id"`
					Precio int `json:"precio"`
				} `json:"precios"`
			} `json:"tipo"`
			Dias []struct {
				Nombre     string `json:"nombre"`
				NumeroDia  int    `json:"numeroDia"`
				GetInicial string `json:"getInicial"`
			} `json:"dias"`
			Nombre             string `json:"nombre"`
			FechaVigenciaDesde struct {
				Hora       string `json:"hora"`
				HoraCorta  string `json:"horaCorta"`
				Fecha      string `json:"fecha"`
				Mysql      string `json:"mysql"`
				DiaNombre  string `json:"diaNombre"`
				FechaMysql string `json:"fechaMysql"`
				Timezone   struct {
					TimezoneType int    `json:"timezone_type"`
					Timezone     string `json:"timezone"`
				} `json:"timezone"`
				Timestamp  int    `json:"timestamp"`
				FechaCorta string `json:"fechaCorta"`
			} `json:"fechaVigenciaDesde"`
			FechaVigenciaHasta interface{} `json:"fechaVigenciaHasta"`
			HoraInicio         struct {
				Hora       string `json:"hora"`
				HoraCorta  string `json:"horaCorta"`
				Fecha      string `json:"fecha"`
				Mysql      string `json:"mysql"`
				DiaNombre  string `json:"diaNombre"`
				FechaMysql string `json:"fechaMysql"`
				Timezone   struct {
					TimezoneType int    `json:"timezone_type"`
					Timezone     string `json:"timezone"`
				} `json:"timezone"`
				Timestamp  int    `json:"timestamp"`
				FechaCorta string `json:"fechaCorta"`
			} `json:"horaInicio"`
			HoraFin struct {
				Hora       string `json:"hora"`
				HoraCorta  string `json:"horaCorta"`
				Fecha      string `json:"fecha"`
				Mysql      string `json:"mysql"`
				DiaNombre  string `json:"diaNombre"`
				FechaMysql string `json:"fechaMysql"`
				Timezone   struct {
					TimezoneType int    `json:"timezone_type"`
					Timezone     string `json:"timezone"`
				} `json:"timezone"`
				Timestamp  int    `json:"timestamp"`
				FechaCorta string `json:"fechaCorta"`
			} `json:"horaFin"`
			SinControlCupos               bool   `json:"sinControlCupos"`
			ParaLlevar                    bool   `json:"paraLlevar"`
			ComprobarTieneFechaHasta      bool   `json:"comprobarTieneFechaHasta"`
			ComprobarEsFutura             bool   `json:"comprobarEsFutura"`
			GetTiempoLimiteReservaMinutos int    `json:"getTiempoLimiteReservaMinutos"`
			Horario                       string `json:"horario"`
		} `json:"servicios"`
	} `json:"serviciosPorTipo"`
}

type ReservarJson struct {
	Comedores []Comedor `json:"comedores"`
	Comensal  struct {
		Facultad struct {
			ID     int    `json:"id"`
			Nombre string `json:"nombre"`
		} `json:"facultad"`
		Carrera struct {
			ID     int    `json:"id"`
			Nombre string `json:"nombre"`
		} `json:"carrera"`
		Confirmado  bool `json:"confirmado"`
		Intercambio bool `json:"intercambio"`
		Categoria   struct {
			ID     int    `json:"id"`
			Nombre string `json:"nombre"`
		} `json:"categoria"`
		PreferenciaPlato struct {
			ID     int    `json:"id"`
			Nombre string `json:"nombre"`
		} `json:"preferenciaPlato"`
		Telefono   string `json:"telefono"`
		Dni        string `json:"dni"`
		Correo     string `json:"correo"`
		Becado     bool   `json:"becado"`
		Registrado bool   `json:"registrado"`
		Apellido   string `json:"apellido"`
		ID         int    `json:"id"`
		Nombre     string `json:"nombre"`
		Foto       string `json:"foto"`
		Habilitado bool   `json:"habilitado"`
		FotoURL    string `json:"fotoUrl"`
		Saldo      int    `json:"saldo"`
	} `json:"comensal"`
	Urls struct {
		GuardarReserva           string `json:"guardarReserva"`
		BuscarTurnosReservas     string `json:"buscarTurnosReservas"`
		DetallePlato             string `json:"detallePlato"`
		ComprobarCancelarReserva string `json:"comprobarCancelarReserva"`
		CancelarReserva          string `json:"cancelarReserva"`
		BuscarReservasComensal   string `json:"buscarReservasComensal"`
	} `json:"urls"`
}
