package internal

import (
	"cmp"
	"flag"
	"os"
)

const (
	defaultHost          = "0.0.0.0"
	defaultPort          = 8080
	defauldDNS           = "postgres://postgres:password@0.0.0.0:5432/postgres?sslmode=disable"
	defaultMigrationPath = "migrations"
)

type Config struct {
	Host        string
	Port        int
	DNS         string
	MigratePath string
}

func ReadConfig() Config {
	var config Config
	flag.StringVar(&config.Host, "host", defaultHost, "адрес для запуска сервера")
	flag.IntVar(&config.Port, "port", defaultPort, "порт для подключения сервера")
	flag.StringVar(&config.DNS, "dns", defauldDNS, "строка подключения к БД")
	flag.StringVar(&config.MigratePath, "migrate-path", defaultMigrationPath, "путь к миграции")
	flag.Parse()

	config.DNS = cmp.Or(os.Getenv("DNS"), defauldDNS)
	config.MigratePath = cmp.Or(os.Getenv("MIGRATE_PATH"), defaultMigrationPath)

	return config
}
