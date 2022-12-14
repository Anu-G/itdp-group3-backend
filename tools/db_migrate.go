package tools

import (
	"itdp-group3-backend/manager"
	"itdp-group3-backend/model/entity"
)

// RunMigrate : tool for db migration
func RunMigrate(dbc manager.InfraManagerInterface) error {
	var err error

	sqlDB, _ := dbc.DBCon().DB()
	defer sqlDB.Close()

	err = dbc.DBCon().AutoMigrate(
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
	if err != nil {
		panic(err)
	}
	return err
}
