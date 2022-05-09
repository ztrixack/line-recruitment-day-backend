package configs

type serverConfig struct {
	Driver  string
	Host    string
	Port    int
	Prefork bool
}

// Server for web server configuration
var Server serverConfig
