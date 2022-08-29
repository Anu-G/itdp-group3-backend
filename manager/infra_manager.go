package manager

import (
	"fmt"
	"itdp-group3-backend/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type InfraManagerInterface interface {
	DBCon() *gorm.DB
	GetMediaPath() string
	GetMediaPathProduct() string
	GetMediaPathFeed() string
	GetMediaPathClientFeed() string
}

type infraManager struct {
	db             *gorm.DB
	path           string
	pathProduct    string
	pathClientFeed string
	pathFeed       string
	cfgDB          config.DBConfig
}

// NewInfraSetup : init new infra manager
func NewInfraSetup(config config.Config) InfraManagerInterface {
	newInfra := new(infraManager)
	newInfra.cfgDB = config.DBConfig
	newInfra.db = newInfra.dbConnect()
	newInfra.path = config.Path
	newInfra.pathProduct = config.PathProduct
	newInfra.pathClientFeed = config.PathClientFeed
	newInfra.pathFeed = config.MediaPath.PathFeed
	return newInfra
}

// dbConnect : connect postgres
func (im *infraManager) dbConnect() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v timezone=%v",
		im.cfgDB.DBHost, im.cfgDB.DBUser, im.cfgDB.DBPassword, im.cfgDB.DBName, im.cfgDB.DBPort, im.cfgDB.SSLMode, im.cfgDB.TimeZone)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	gormCfg := new(gorm.Config)
	if im.cfgDB.Environment == "DEV" {
		gormCfg.Logger = newLogger
	}

	dbcon, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, _ := dbcon.DB()
	defer sqlDB.SetConnMaxLifetime(15 * time.Minute)

	return dbcon
}

// DBCon : return gorm database
func (im *infraManager) DBCon() *gorm.DB {
	return im.db
}

func (im *infraManager) GetMediaPath() string {
	return im.path
}

func (im *infraManager) GetMediaPathProduct() string {
	return im.pathProduct
}

func (im *infraManager) GetMediaPathFeed() string {
	return im.pathFeed
}

func (im *infraManager) GetMediaPathClientFeed() string {
	return im.pathClientFeed
}
