package main

import (
	"fmt"
	"github.com/fy23-gw-gackathon/reportify-backend/config"
	"github.com/fy23-gw-gackathon/reportify-backend/controller"
	"github.com/fy23-gw-gackathon/reportify-backend/docs"
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
	api := app.Group("/api/v1")
	api.GET("/users/me", handleResponse(userController.GetMe))
	api.PUT("/reports/:reportId", handleResponse(reportController.ReviewReport, http.StatusNoContent))
	orgs := api.Group("/organizations")
	orgs.Use(middleware.Authentication(userPersistence, cfg))
	{
		orgs.GET("/", handleResponse(organizationController.GetOrganizations))
		orgs.GET("/:organizationCode", handleResponse(organizationController.GetOrganization))
		orgs.PUT("/:organizationCode", handleResponse(organizationController.UpdateOrganization))

		orgs.GET("/:organizationCode/reports", handleResponse(reportController.GetReports))
		orgs.POST("/:organizationCode/reports", handleResponse(reportController.CreateReport, http.StatusCreated))
		orgs.GET("/:organizationCode/reports/:reportId", handleResponse(reportController.GetReport))

		orgs.GET("/:organizationCode/users", handleResponse(userController.GetUsers))
		orgs.POST("/:organizationCode/users", handleResponse(userController.InviteUser))
		orgs.PUT("/:organizationCode/users/:userId", handleResponse(userController.UpdateUserRole))
		orgs.DELETE("/:organizationCode/users/:userId", handleResponse(userController.DeleteUser))
	}

	runApp(app, cfg.App.Port)
}

func runApp(app *gin.Engine, port int) {
	docs.SwaggerInfo.Title = "Reportify"
	docs.SwaggerInfo.Description = "Reportify"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
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
