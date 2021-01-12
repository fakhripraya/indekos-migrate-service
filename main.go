package main

import (
	"log"
	"strconv"

	"github.com/fakhripraya/indekos-migrate-service/config"
	"github.com/fakhripraya/indekos-migrate-service/data"
	"github.com/fakhripraya/indekos-migrate-service/entities"
	"github.com/fakhripraya/indekos-migrate-service/migrate"
	"github.com/hashicorp/go-hclog"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

func main() {

	// creates a structured logger for logging the entire program
	logger := hclog.Default()

	// load configuration from env file
	err = godotenv.Load(".env")

	if err != nil {
		// log the fatal error if load env failed
		log.Fatal(err)
	}

	// Initialize app configuration
	var appConfig entities.Configuration
	data.ConfigInit(&appConfig)

	// Open the database connection based on DB configuration
	logger.Info("Establishing database connection on " + appConfig.Database.Host + ":" + strconv.Itoa(appConfig.Database.Port))
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig(&appConfig.Database)))
	if err != nil {
		log.Fatal(err)
	}

	defer config.DB.Close()

	config.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&migrate.MasterUser{},
		&migrate.MasterUserLogin{},
		&migrate.MasterAccess{},
		&migrate.MasterRole{},
		&migrate.MasterKostType{},
		&migrate.MasterFacilities{},
		&migrate.DBKost{},
		&migrate.DBKostRoom{},
		&migrate.DBKostFacilities{},
	)
}
