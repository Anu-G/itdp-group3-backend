package tools

import "itdp-group3-backend/manager"

// RunMigrate : tool for db migration
func RunMigrate(dbc manager.InfraManagerInterface) error {
	var err error

	sqlDB, _ := dbc.DBCon().DB()
	defer sqlDB.Close()

	err = dbc.DBCon().AutoMigrate(
	// put entity models here
	)
	if err != nil {
		panic(err)
	}
	return err
}
