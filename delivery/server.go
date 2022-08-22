package delivery

import (
	"itdp-group3-backend/config"
	"itdp-group3-backend/delivery/controller"
	"itdp-group3-backend/manager"
	"itdp-group3-backend/tools"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	engine      *gin.Engine
	host        string
	startServer bool

	UseCaseManager manager.UseCaseManagerInterface
}

// Server : prepare config and read arguments
func Server() *appServer {
	r := gin.Default()

	appCfg := config.NewConfig()
	dbCon := manager.NewInfraSetup(appCfg)
	repoManager := manager.NewRepo(dbCon)
	usecaseManager := manager.NewUseCase(repoManager)

	cfgServer := &appServer{
		engine:         r,
		host:           appCfg.APIConfig.APIUrl,
		UseCaseManager: usecaseManager,
	}

	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "db:migrate":
			if appCfg.DBConfig.Environment == "DEV" {
				tools.RunMigrate(dbCon)
			} else {
				log.Fatal("not a dev env, cannot migrate")
			}
			return cfgServer
		case "db:seeds":
			if appCfg.DBConfig.Environment == "DEV" {
				tools.RunSeed(dbCon)
			} else {
				log.Fatal("not a dev env, cannot seeding")
			}
			return cfgServer
		default:
			log.Fatalln("argument not found !")
			cfgServer.startServer = true
			return cfgServer
		}
	}

	cfgServer.startServer = true
	return cfgServer
}

// initControllers : prepare the controller API
func (a *appServer) initControllers() {
	controller.NewBusinessProfileController(a.engine, a.UseCaseManager.BusinessProfileUseCase())
	controller.NewNonBusinessProfileController(a.engine, a.UseCaseManager.NonBusinessProfileUseCase())
	controller.NewProductController(a.engine,a.UseCaseManager.ProductUseCase() )
}

// Run : run the server
func (a *appServer) Run() {
	if a.startServer {
		a.initControllers()
		if err := a.engine.Run(a.host); err != nil {
			panic(err)
		}
	}
}
