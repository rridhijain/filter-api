package config

// Configurations exported
type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	Host     string
	Port     int
	DBName string
	User string
	Password string
}