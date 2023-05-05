package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"reportify-backend/controller"
	"reportify-backend/docs"
	"reportify-backend/entity"
	"reportify-backend/infrastructure/driver"
	"reportify-backend/infrastructure/middleware"
	"reportify-backend/infrastructure/persistence"
	"reportify-backend/usecase"
)

const (
	defaultPort = "8080"
)

func main() {
	// Dependency Injection
	db := driver.NewDB()

	userPersistence := persistence.NewUserPersistence()
	organizationPersistence := persistence.NewOrganizationPersistence()

	userUseCase := usecase.NewUserUseCase(userPersistence)
	organizationUseCase := usecase.NewOrganizationUseCase(organizationPersistence)

	userController := controller.NewUserController(userUseCase)
	organizationController := controller.NewOrganizationController(organizationUseCase)
	reportController := controller.NewReportController(nil)

	// Setup webserver
	app := gin.Default()

	app.Use(middleware.Transaction(db))
	app.Use(middleware.Cors())
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "It works")
	})
	app.GET("/organizations", handleResponse(organizationController.GetOrganizations))

	org := app.Group("/organizations/:organizationCode")
	org.GET("/", handleResponse(organizationController.GetOrganization))
	org.PUT("/", handleResponse(organizationController.UpdateOrganization))

	org.GET("/reports", handleResponse(reportController.GetReports))
	org.POST("/reports", handleResponse(reportController.CreateReport))
	org.GET("/reports/:reportId", handleResponse(reportController.GetReport))

	org.GET("/users", handleResponse(userController.GetUsers))

	runApp(app)
}

func runApp(app *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	docs.SwaggerInfo.Title = "Reportify"
	docs.SwaggerInfo.Description = "Reportify"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", port)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println(fmt.Sprintf("http://localhost:%s", port))
	log.Println(fmt.Sprintf("http://localhost:%s/swagger/index.html", port))
	app.Run(fmt.Sprintf(":%s", port))
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
