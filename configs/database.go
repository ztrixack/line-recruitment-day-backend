package configs

type databaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string
}

// Database for database configuration
var Database databaseConfig
