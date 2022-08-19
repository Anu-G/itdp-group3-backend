package delivery

import (
	"itdp-group3-backend/auth"
	"itdp-group3-backend/config"
	"itdp-group3-backend/controller"
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
	Auth        auth.Token

	UseCaseManager    manager.UseCaseManagerInterface
	MiddlewareManager manager.MiddlewareManager
}

// Server : prepare config and read arguments
func Server() *appServer {
	r := gin.Default()

	appCfg := config.NewConfig()
	dbCon := manager.NewInfraSetup(appCfg)
	auth := auth.NewTokenService(appCfg.TokenConfig)
	repoManager := manager.NewRepo(dbCon)
	usecaseManager := manager.NewUseCase(repoManager)
	middlewareManager := manager.NewMiddlewareManager(auth)

	cfgServer := &appServer{
		engine:            r,
		host:              appCfg.APIConfig.APIUrl,
		Auth:              auth,
		UseCaseManager:    usecaseManager,
		MiddlewareManager: middlewareManager,
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
	controller.NewAccountController(a.engine, a.UseCaseManager.AccountUsecase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewAuthController(a.engine, a.UseCaseManager.AuthUsecase(), a.UseCaseManager.UserUsecase(), a.Auth)
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
