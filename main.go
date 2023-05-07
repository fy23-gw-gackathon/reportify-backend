package main

import (
	"fmt"
	"github.com/fy23-gw-gackathon/reportify-backend/config"
	"github.com/fy23-gw-gackathon/reportify-backend/controller"
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/driver"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/middleware"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/persistence"
	"github.com/fy23-gw-gackathon/reportify-backend/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

// @title                      Reportify
// @version                    1.0
// @description                Reportify
// @host                       localhost:8080
// @BasePath                   /
// @schemes                    http
// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and JWT token.
func main() {
	// Dependency Injection
	cfg := config.Load()
	db := driver.NewDB(cfg)
	redis := driver.NewRedisClient(cfg.Datastore.Address)
	cognitoClient := driver.NewCognitoClient(cfg)

	userPersistence := persistence.NewUserPersistence(cognitoClient)
	organizationPersistence := persistence.NewOrganizationPersistence()
	reportPersistence := persistence.NewReportPersistence(redis)

	userUseCase := usecase.NewUserUseCase(userPersistence, organizationPersistence)
	organizationUseCase := usecase.NewOrganizationUseCase(organizationPersistence, userPersistence)
	reportUseCase := usecase.NewReportUseCase(reportPersistence, userPersistence, organizationPersistence)

	userController := controller.NewUserController(userUseCase)
	organizationController := controller.NewOrganizationController(organizationUseCase)
	reportController := controller.NewReportController(reportUseCase)

	// Setup webserver
	app := gin.Default()

	app.Use(middleware.Transaction(db))
	app.Use(middleware.Cors(cfg))
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "It works")
	})
	app.GET("/users/me", handleResponse(userController.GetMe))
	app.PUT("/reports/:reportId", handleResponse(reportController.ReviewReport, http.StatusNoContent))

	orgs := app.Group("/organizations")
	orgs.Use(middleware.Authentication(userPersistence, cfg))
	orgs.GET("/", handleResponse(organizationController.GetOrganizations))
	org := orgs.Group("/:organizationCode")
	org.GET("/", handleResponse(organizationController.GetOrganization))
	org.PUT("/", handleResponse(organizationController.UpdateOrganization))

	org.GET("/reports", handleResponse(reportController.GetReports))
	org.POST("/reports", handleResponse(reportController.CreateReport, http.StatusCreated))
	org.GET("/reports/:reportId", handleResponse(reportController.GetReport))

	org.GET("/users", handleResponse(userController.GetUsers))
	org.POST("/users", handleResponse(userController.InviteUser))

	runApp(app, cfg.App.Port)
}

func runApp(app *gin.Engine, port int) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println(fmt.Sprintf("http://localhost:%d", port))
	log.Println(fmt.Sprintf("http://localhost:%d/swagger/index.html", port))
	app.Run(fmt.Sprintf(":%d", port))
}

func handleResponse(f func(ctx *gin.Context) (interface{}, error), status ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := f(c)
		if err != nil {
			e, ok := err.(*entity.Error)
			if ok {
				c.JSON(e.Code, entity.ErrorResponse{Message: err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, entity.ErrorResponse{Message: err.Error()})
			}
			c.Abort()
		} else {
			if len(status) > 0 {
				c.JSON(status[0], result)
			} else {
				c.JSON(http.StatusOK, result)
			}
		}
	}
}
