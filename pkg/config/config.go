package config

type Config struct {
	Server        Server
	MocksFilePath string
}

type Server struct {
	Port string
}

var Cfg = Config{
	Server: Server{
		Port: "8080",
	},
	MocksFilePath: "./default.json",
}
