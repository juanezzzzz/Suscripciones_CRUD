package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// ConnectDB opens a database connection and returns any error.
func ConnectDB() error {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "3105437922")
	dbname := getEnv("DB_NAME", "Jobsy")
	schema := getEnv("DB_SCHEMA", "suscripciones")

	psqlinfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema,
	)

	var err error
	DB, err = sql.Open("postgres", psqlinfo)
	if err != nil {
		return fmt.Errorf("error al abrir la conexión a la base de datos: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error al hacer ping a la base de datos: %w", err)
	}

	log.Println("Conectado a la base de datos")
	log.Println("Schema:", schema)
	return nil
}
