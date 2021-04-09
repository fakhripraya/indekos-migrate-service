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
	"github.com/srinathgs/mysqlstore"
)

var err error

// Session Store based on MYSQL database
var sessionStore *mysqlstore.MySQLStore

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
	err = data.ConfigInit(&appConfig)

	if err != nil {
		// log the fatal error if config init failed
		log.Fatal(err)
	}

	// Open the database connection based on DB configuration
	logger.Info("Establishing database connection on " + appConfig.Database.Host + ":" + strconv.Itoa(appConfig.Database.Port))
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig(&appConfig.Database)))
	if err != nil {
		log.Fatal(err)
	}

	defer config.DB.Close()

	err = config.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&migrate.MasterUser{},
		&migrate.MasterRole{},
		&migrate.MasterUOM{},
		&migrate.MasterAccess{},
		&migrate.MasterPeriod{},
		&migrate.MasterUserLogin{},
		&migrate.MasterKostType{},
		&migrate.MasterRoleAccess{},
		&migrate.MasterFacilities{},
		&migrate.MasterPaymentMethod{},
		&migrate.MasterEvent{},
		&migrate.MasterEventDetail{},
		&migrate.MasterIcon{},
		&migrate.DBKost{},
		&migrate.DBKostPeriod{},
		&migrate.DBKostPict{},
		&migrate.DBKostFacilities{},
		&migrate.DBKostReview{},
		&migrate.DBKostBenchmark{},
		&migrate.DBKostAccess{},
		&migrate.DBKostAround{},
		&migrate.DBKostRoom{},
		&migrate.DBKostRoomDetail{},
		&migrate.DBKostRoomPict{},
		&migrate.DBKostRoomFacilities{},
		&migrate.DBTransaction{},
		&migrate.DBTransactionDetail{},
		&migrate.DBTransactionVerification{},
		&migrate.DBTransactionRoomBook{},
		&migrate.DBTransactionRoomBookMember{},
	).Error

	if err != nil {
		log.Fatal(err)
	}
}
