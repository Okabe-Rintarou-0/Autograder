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
	groupDAO := dao.NewGroupDAO(systemDB, eBookStoreDB)
	groupSvc := service.NewGroupService(groupDAO)
	h := handler.NewHandler(groupSvc)
	r.Group("/api", interceptor.NewTokenInterceptor(cfg.Token)).
		GET("/courses", h.HandleGetCourses).
		GET("/me", h.HandleGetMe).
		GET("/tasks", h.HandleListAppTasks).
		GET("/users", h.HandleListUsers).
		POST("/run", h.HandleRunApp).
		POST("/register", h.HandleRegister).
		POST("/register/canvas", h.HandleImportCanvasUsers).
		PUT("/me/password", h.HandleChangePassword)

	r.POST("/api/login", h.HandleLogin).
		GET("/api/logs", h.HandleGetLog)
	logrus.Fatal(r.Run(":8081"))
}
