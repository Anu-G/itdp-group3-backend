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
		&entity.UserCredential{},
	)
	if err != nil {
		panic(err)
	}
	return err
}
