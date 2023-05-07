package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"reportify-backend/config"
	"reportify-backend/controller"
	"reportify-backend/docs"
	"reportify-backend/entity"
	"reportify-backend/infrastructure/driver"
	"reportify-backend/infrastructure/middleware"
	"reportify-backend/infrastructure/persistence"
	"reportify-backend/usecase"
)

// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and JWT token.
func main() {
	// Dependency Injection
	cfg := config.Load()
	db := driver.NewDB(cfg)
	cognitoClient := driver.NewCognitoClient(cfg)

	userPersistence := persistence.NewUserPersistence(cognitoClient)
	organizationPersistence := persistence.NewOrganizationPersistence()
	reportPersistence := persistence.NewReportPersistence()

	userUseCase := usecase.NewUserUseCase(userPersistence, organizationPersistence)
	organizationUseCase := usecase.NewOrganizationUseCase(organizationPersistence, userPersistence)
	reportUseCase := usecase.NewReportUseCase(reportPersistence, userPersistence)

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
	orgs := app.Group("/organizations")
	orgs.Use(middleware.Authentication(userPersistence))
	orgs.GET("/", handleResponse(organizationController.GetOrganizations))
	org := orgs.Group("/:organizationCode")
	org.GET("/", handleResponse(organizationController.GetOrganization))
	org.PUT("/", handleResponse(organizationController.UpdateOrganization))

	org.GET("/reports", handleResponse(reportController.GetReports))
	org.POST("/reports", handleResponse(reportController.CreateReport))
	org.GET("/reports/:reportId", handleResponse(reportController.GetReport))

	org.GET("/users", handleResponse(userController.GetUsers))
	org.POST("/users", handleResponse(userController.InviteUser))

	app.GET("/users/me", handleResponse(userController.GetMe))

	runApp(app, cfg.App.Port)
}

func runApp(app *gin.Engine, port int) {
	docs.SwaggerInfo.Title = "Reportify"
	docs.SwaggerInfo.Description = "Reportify"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", port)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println(fmt.Sprintf("http://localhost:%d", port))
	log.Println(fmt.Sprintf("http://localhost:%d/swagger/index.html", port))
	app.Run(fmt.Sprintf(":%d", port))
}

func handleResponse(f func(ctx *gin.Context) (interface{}, error)) gin.HandlerFunc {
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
			c.JSON(http.StatusOK, result)
		}
	}
}
