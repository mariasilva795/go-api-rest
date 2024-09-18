package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
Creación del servidor:

1. La función NewServer() valida la configuración y devuelve una instancia del servidor (Broker).
Configuración de rutas:

2. El método Start() acepta una función binder que se utiliza para vincular rutas al enrutador (mux.Router).
Inicio del servidor:

3. El servidor comienza a escuchar en el puerto especificado y maneja las solicitudes usando el enrutador que se configuró.

*/

// Config: tiene las caracteristicas del servidor. El puerto en el que se va ejecutar, la clave secreta para generar tokens y la conexion a base de datos.
type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

// Server: esta interface implementa el modelo de datos o estructura de config.
type Server interface {
	Config() *Config
}

// Broker: Nos ayuda a tener varias instancias de servidor corriendo. Esta estructura a su vez tiene la estructura Config y el metodo Config, para ser de tipo Server.
type Broker struct {
	config *Config
	router *mux.Router
}

// Methos in go
func (b *Broker) Config() *Config {
	return b.config
}

// Validaciones de la entrada
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("jwtsecret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	log.Println("Starting our server", b.Config().Port)

	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
