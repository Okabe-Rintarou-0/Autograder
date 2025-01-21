package main

import (
	"autograder/pkg/config"
	"autograder/pkg/dal/mysql"
	"autograder/pkg/dao"
	"autograder/pkg/handler"
	"autograder/pkg/interceptor"
	"autograder/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// ctx := context.Background()
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	r := gin.Default()
	cfg := config.Instance
	systemDB, err := mysql.NewDB(cfg.SystemDB)
	if err != nil {
		panic(err)
	}
	eBookStoreDB, err := mysql.NewDB(cfg.EBookStoreDB)
	if err != nil {
		panic(err)
	}
	dao := dao.NewGroupDAO(systemDB, eBookStoreDB)
	groupSvc := service.NewGroupService(dao)
	handler := handler.NewHandler(groupSvc)
	r.Group("/api", interceptor.NewTokenInterceptor(cfg.Token)).
		GET("/me", handler.HandleGetMe).
		GET("/tasks", handler.HandleListAppTasks).
		POST("/run", handler.HandleRunApp).
		PUT("/me/password", handler.HandleChangePassword)
	r.POST("/api/login", handler.HandleLogin).
		GET("/api/logs", handler.HandleGetLog)
	logrus.Fatal(r.Run(":8081"))
}
