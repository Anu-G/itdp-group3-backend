package manager

import (
	"database/sql"
	"itdp-group3-backend/config"
	"itdp-group3-backend/model/entity"
	"log"
	"os"
	"time"

	"github.com/lib/pq"
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

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("pgx", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbcon, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), gormCfg)
	if err != nil {
		log.Fatal(err)
	}

	if !dbcon.Migrator().HasTable(&entity.User{}) || !dbcon.Migrator().HasTable(&entity.Account{}) || !dbcon.Migrator().HasTable(&entity.Category{}) || !dbcon.Migrator().HasTable(&entity.BusinessProfile{}) || !dbcon.Migrator().HasTable(&entity.NonBusinessProfile{}) || !dbcon.Migrator().HasTable(&entity.Product{}) || !dbcon.Migrator().HasTable(&entity.Feed{}) || !dbcon.Migrator().HasTable(&entity.BusinessFAQ{}) || !dbcon.Migrator().HasTable(&entity.BusinessHour{}) || !dbcon.Migrator().HasTable(&entity.BusinessLink{}) || !dbcon.Migrator().HasTable(&entity.DetailMediaFeed{}) || !dbcon.Migrator().HasTable(&entity.DetailComment{}) || !dbcon.Migrator().HasTable(&entity.Followed{}) || !dbcon.Migrator().HasTable(&entity.Follower{}) {
		dbcon.AutoMigrate(
			// put entity models here
			&entity.User{},
			&entity.Account{},
			&entity.Category{},
			&entity.BusinessProfile{},
			&entity.NonBusinessProfile{},
			&entity.Product{},
			&entity.Feed{},
			&entity.BusinessFAQ{},
			&entity.BusinessHour{},
			&entity.BusinessLink{},
			&entity.DetailMediaFeed{},
			&entity.DetailComment{},
			&entity.Followed{},
			&entity.Follower{},
		)
	}

	db.SetConnMaxLifetime(15 * time.Minute)

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
