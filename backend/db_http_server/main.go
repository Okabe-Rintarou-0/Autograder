package db_http_server

import (
	"fmt"
	"net/http"

	"autograder/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// example:
/*
curl -X POST \
http://localhost:5000/execute \
-H 'Content-Type: application/json' \
-d '{
  "database": "ebookstore",
  "sql": "select * from users"
}'
*/

const (
	username = "root"
	password = "123456"
	hostname = "127.0.0.1:3306"
)

type executeParam struct {
	Database string `json:"database" form:"database"`
	SQL      string `json:"sql" form:"sql"`
}

func executeSQL(c *gin.Context) {
	param := executeParam{}
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "database and sql parameters are required",
		})
		return
	}
	database := param.Database
	sql := param.SQL

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to connect to database: %v", err),
		})
		return
	}

	var result []map[string]any
	err = db.Raw(sql).Scan(&result).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to execute SQL: %v", err),
		})
		return
	}

	logrus.Infof("result: %s", utils.FormatJsonString(result))

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func RunDBHTTPServer() {
	r := gin.Default()
	r.POST("/execute", executeSQL)
	logrus.Fatal(r.Run(":5000"))
}
