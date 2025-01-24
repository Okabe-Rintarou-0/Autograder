package main

import (
	"autograder/db_http_server"
	"autograder/pkg/config"
	"autograder/pkg/dal/mysql"
	"autograder/pkg/dao"
	"autograder/pkg/handler"
	"autograder/pkg/interceptor"
	"autograder/pkg/model/dbm"
	"autograder/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// ctx := context.Background()
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	go db_http_server.RunDBHTTPServer()

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
		GET("/me", h.HandleGetMe).
		GET("/tasks", h.HandleListAppTasks).
		GET("/users", h.HandleListUsers).
		POST("/run", h.HandleRunApp).
		PUT("/me/password", h.HandleChangePassword)

	r.Group("/api",
		interceptor.NewTokenInterceptor(cfg.Token),
		interceptor.NewRoleInterceptor(dbm.Administrator, groupDAO)).
		GET("/courses", h.HandleGetCourses).
		GET("/assignments", h.HandleGetAssignments).
		GET("/submissions", h.HandleGetAssignmentSubmissions).
		POST("/register", h.HandleRegister).
		POST("/register/canvas", h.HandleImportCanvasUsers).
		GET("/canvas/users", h.HandleGetCourseUsers)

	r.POST("/api/login", h.HandleLogin).
		GET("/api/logs", h.HandleGetLog)
	logrus.Fatal(r.Run(":8081"))
}
