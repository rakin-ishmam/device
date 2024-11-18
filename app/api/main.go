package main

import (
	"device/app/api/handler"
	"device/business/device"
	"device/config"
	"fmt"
	"net/http"
	"time"

	deviceHandler "device/app/api/handler/device"
	deviceStore "device/business/device/store/postgres"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	appCfg := config.MustLoad()

	db := initDBClient(appCfg.Database)

	deviceStore := deviceStore.NewStore(db)
	deviceBusiness := device.NewBusiness(deviceStore)
	router := handler.NewRouter(handler.Handlers{
		Device: deviceHandler.NewHandler(deviceBusiness),
	})

	log.Info("Starting server on port ", appCfg.Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", appCfg.Server.Port), router); err != nil {
		log.Fatalf("unable to start server: %v", err)
	}
}

func initDBClient(cfg config.Database) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s TimeZone=UTC",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	// Initialize GORM with the PostgreSQL driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log SQL queries for debugging
	})
	if err != nil {
		panic(fmt.Errorf("failed to connect to PostgreSQL: %w", err))
	}

	// Configure the connection pool (optional)
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("failed to get database handle: %w", err))
	}

	sqlDB.SetMaxOpenConns(25)                 // Maximum number of open connections
	sqlDB.SetMaxIdleConns(10)                 // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime

	log.Info("PostgreSQL client initialized successfully")

	// migration
	if err := db.AutoMigrate(&deviceStore.Device{}); err != nil {
		panic(fmt.Errorf("failed to migrate database: %w", err))
	}

	return db
}

func initHandlers() {

}

func startServer() {

}
