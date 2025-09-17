package internal

import "flag"

type Config struct {
	Host string
	Port int
}

func ReadConfig() Config {
	var config Config
	flag.StringVar(&config.Host, "host", "localhost", "адрес для запуска сервера")
	flag.IntVar(&config.Port, "port", 8080, "порт для подключения сервера")
	flag.Parse()
	return config
}
