package config

import(
	"log"
	"os"
)

var env string

func init() {
	env = os.Getenv("ENV")
	if env == "" {
		log.Fatal("ENV variable is empty")
	}
}