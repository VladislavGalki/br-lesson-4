package internal

import "flag"

type Config struct {
	Host        string
	Port        int
	DSN         string
	MigratePath string
}

func ReadConfig() Config {
	var config Config
	flag.StringVar(&config.Host, "host", "localhost", "адрес для запуска сервера")
	flag.IntVar(&config.Port, "port", 8080, "порт для подключения сервера")
	flag.StringVar(&config.DSN, "dsn", "postgres://postgres:password@0.0.0.0:5432/postgres?sslmode=disable", "строка подключения к БД")
	flag.StringVar(&config.MigratePath, "migrate-path", "migrations", "путь к миграции")
	flag.Parse()
	return config
}
