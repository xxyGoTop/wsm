package config

import "github.com/xxyGoTop/wsm/internal/lib/dotenv"

type database struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	Driver       string `json:"drver"`
	DatabaseName string `json:"database_name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Sync         string `json:"sync"`
}

// Database message
var Database database

func init() {
	Database.Driver = dotenv.GetByDefault("DB_Driver", "postgres")
	Database.Host = dotenv.GetByDefault("DB_Host", "localhost")
	Database.Port = dotenv.GetByDefault("DB_PORT", "54321")
	Database.DatabaseName = dotenv.GetByDefault("DB_NAME", "terminal")
	Database.Username = dotenv.GetByDefault("DB_USERNAME", "xxy")
	Database.Password = dotenv.GetByDefault("DB_PASSWORD", "xxy")
	Database.Sync = dotenv.GetByDefault("DB_SYNC", "on")
}
