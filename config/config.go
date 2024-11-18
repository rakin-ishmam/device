package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// App represents the application configuration
type App struct {
	Server   Server
	Database Database
}

// Load loads the application configuration
func (a *App) Load() {
	a.Server.Load()
	a.Database.Load()
}

// Server represents the server configuration
type Server struct {
	Port string
}

// Load loads the server configuration
func (s *Server) Load() {
	s.Port = os.Getenv("SERVER_PORT")
}

// Database represents the database configuration
type Database struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

// Load loads the database configuration
func (d *Database) Load() {
	d.Host = os.Getenv("DB_HOST")
	d.Port = os.Getenv("DB_PORT")
	d.DBName = os.Getenv("DB_NAME")
	d.User = os.Getenv("DB_USER")
	d.Password = os.Getenv("DB_PASSWORD")
}

// MustLoad loads the application configuration and panics if an error occurs
func MustLoad() App {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := App{}
	app.Load()
	return app
}
