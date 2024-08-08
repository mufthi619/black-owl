package config

type AppConfig struct {
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     string `json:"port"`
}
