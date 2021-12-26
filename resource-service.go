package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"resource_service/internal/adaptars/api/resource"
	"resource_service/internal/adaptars/repositories/mongodb/resource"
	"resource_service/internal/core/helper"
	"resource_service/internal/core/services/resource"
	"resource_service/internal/core/shared"
	port "resource_service/internal/ports/resource"
	"time"
)

func main() {
	helper.InitializeLog()
	address, prt, _, dbHost, dbname, _ := helper.LoadConfig()
	dbRepository := ConnectToMongo(dbHost, dbname)
	service := services.New(dbRepository)
	handler := api.NewHTTPHandler(service)
	router := gin.Default()
	router.Use(helper.LogRequest)
	router.GET("/resource/:reference", handler.Read)
	router.GET("/resource/entries", handler.ReadAll)
	router.POST("/resource", handler.Create)
	router.PUT("/resource/:reference", handler.Update)
	router.DELETE("/resource/:reference", handler.Delete)
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,
			helper.PrintErrorMessage("404", shared.NoResourceFound))
	})
	fmt.Println("Service running on " + address + ":" + prt)
	helper.LogEvent("Info", fmt.Sprintf("Started PlatformServiceApplication on "+address+":"+prt+" in "+time.Since(time.Now()).String()))
	_ = router.Run(":" + prt)
}
func ConnectToMongo(dbHost string, dbname string) port.ResourceRepository {
	mongoUrl := dbHost + "/?directConnection=true&serverSelectionTimeoutMS=2000"
	repo, err := repository.NewMongoRepository(mongoUrl, dbname, 2000)
	if err != nil {
		_ = helper.PrintErrorMessage("500", err.Error())
		log.Fatal(err)

	}
	return services.New(repo)
}
