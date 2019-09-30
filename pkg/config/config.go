package config

type Config struct {
	Server Server
}

type Server struct {
	Port string
}

var Cfg = Config{
	Server: Server{
		Port: "8080",
	},
}
