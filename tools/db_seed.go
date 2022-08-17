package tools

import "itdp-group3-backend/manager"

// RunSeed : tool for db seeding
func RunSeed(dbc manager.InfraManagerInterface) {
	sqlDB, _ := dbc.DBCon().DB()
	defer sqlDB.Close()

	// repoMng := managers.NewRepo(dbc)

	// put database seeder(init dummy data) here
}
