package delivery

import (
	"itdp-group3-backend/auth"
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
	Auth        auth.Token

	UseCaseManager    manager.UseCaseManagerInterface
	MiddlewareManager manager.MiddlewareManager
}

// Server : prepare config and read arguments
func Server() *appServer {
	r := gin.Default()
	r.Use(CORSMiddleware())

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

	// first time migration for production
	if appCfg.DBConfig.Environment == "PROD" {
		tools.RunMigrate(dbCon)
	} else {
		log.Fatal("cannot migrate")
	}

	cfgServer.startServer = true
	return cfgServer
}

// initControllers : prepare the controller API
func (a *appServer) initControllers() {
	controller.NewUserController(a.engine, a.UseCaseManager.UserUsecase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewAccountController(a.engine, a.UseCaseManager.AccountUsecase(), a.MiddlewareManager.AuthMiddleware(), a.UseCaseManager.FollowUsecase())
	controller.NewDetailMediaFeedController(a.engine, a.UseCaseManager.DetailMediaFeedUsecase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewFeedController(a.engine, a.UseCaseManager.FeedUsecase(), a.UseCaseManager.DetailMediaFeedUsecase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewCommentController(a.engine, a.UseCaseManager.DetailCommentUsecase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewAuthController(a.engine, a.UseCaseManager.AuthUsecase(), a.UseCaseManager.UserUsecase(), a.Auth)
	controller.NewProductController(a.engine, a.UseCaseManager.ProductUseCase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewBusinessProfileController(a.engine, a.UseCaseManager.BusinessProfileUseCase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewNonBusinessProfileController(a.engine, a.UseCaseManager.NonBusinessProfileUseCase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewCategoryController(a.engine, a.UseCaseManager.CategoryUsecase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewFAQController(a.engine, a.UseCaseManager.FAQUseCase(), a.MiddlewareManager.AuthMiddleware())
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
