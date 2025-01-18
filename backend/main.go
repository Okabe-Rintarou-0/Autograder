package main

import (
	"autograder/pkg/dao"
	"autograder/pkg/handler"
	apprunner "autograder/pkg/service/app_runner"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// ctx := context.Background()
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	// info := &entity.AppInfo{
	// 	ZipFileName:        "backend.zip",
	// 	StudentID:          "lucas",
	// 	UploadTime:         time.Now(),
	// 	AuthenticationType: entity.ByCookies,
	// 	JDKVersion:         17,
	// }
	// service := apprunner.NewService(dao.NewGroupDAO())
	// err := service.RunApp(ctx, info)
	// if err != nil {
	// 	logrus.Errorf("err: %+v", err)
	// }
	r := gin.Default()
	dao := dao.NewGroupDAO()
	handler := handler.NewHandler(apprunner.NewService(dao))
	r.POST("/api/run", handler.HandleRunApp)
	r.Run(":8081")
}
