package handler

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"autograder/pkg/config"
	"autograder/pkg/entity"
	apprunner "autograder/pkg/service/app_runner"
)

type Handler struct {
	appRunnerSvc apprunner.Service
}

func NewHandler(appRunnerSvc apprunner.Service) *Handler {
	return &Handler{appRunnerSvc}
}

func getFormIntAttr(c *gin.Context, key string) (int64, error) {
	attrStr := c.Request.Form.Get(key)
	attr, err := strconv.ParseInt(attrStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return attr, nil
}

func (h *Handler) validateParams(c *gin.Context) (*entity.AppInfo, error) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from request"})
		return nil, err
	}

	fileExt := filepath.Ext(file.Filename)
	if fileExt != ".zip" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type, only zip files are allowed"})
		return nil, err
	}

	savePath := path.Join(config.WorkDir, file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return nil, err
	}
	jdkVersion, err := getFormIntAttr(c, "jdk_version")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param, should container form key 'jdk_version'"})
		return nil, err
	}
	authenticationType, err := getFormIntAttr(c, "authentication_type")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param, should container form key 'authentication_type'"})
		return nil, err
	}

	info := &entity.AppInfo{
		ZipFileName:        file.Filename,
		UploadTime:         time.Now(),
		AuthenticationType: entity.AuthenticationType(authenticationType),
		JDKVersion:         int32(jdkVersion),
	}

	if !info.Validate() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param"})
		return nil, fmt.Errorf("invalid info")
	}

	return info, nil
}

func (h *Handler) HandleRunApp(c *gin.Context) {
	appInfo, err := h.validateParams(c)
	if appInfo == nil || err != nil {
		return
	}

	h.appRunnerSvc.RunApp(c.Request.Context(), appInfo)

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
