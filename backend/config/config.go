package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Representa toda la configuración de la aplicación.
type Config struct {
	DatabaseURL string
	ServerPort  string
	AppEnv      string
}

// Variable global de la configuración de la aplicación.
var App Config

// Solo se ejecutará una vez cuando se inicie la aplicación.
// La única responsabilidad de esta función es:
// - Leer las variables de entorno
// - Obtener las variables de entorno
// - Llenar la variable "config.App" (la global declarada aquí arriba)
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
// Después de esto, la aplicación nunca tendrá que volver a leer directamente 'os.GetEnv()'
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
func Load() {

	err := godotenv.Load()

	if err != nil {
		log.Println("[CONFIG] .env file not found, using system environment")
	}

	App = Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ServerPort:  os.Getenv("SERVER_PORT"),
		AppEnv:      os.Getenv("APP_ENV"),
	}
}
