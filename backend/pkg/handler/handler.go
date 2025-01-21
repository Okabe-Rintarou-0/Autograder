package handler

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"autograder/pkg/model/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"autograder/pkg/config"
	"autograder/pkg/messages"
	"autograder/pkg/model/entity"
	"autograder/pkg/model/request"
	"autograder/pkg/model/response"
	"autograder/pkg/service"
	"autograder/pkg/utils"
)

type Handler struct {
	groupSvc *service.GroupService
}

const (
	DefaultPageNo   = 0
	DefaultPageSize = 10
)

func NewHandler(groupSvc *service.GroupService) *Handler {
	return &Handler{groupSvc}
}

func getFormIntAttr(c *gin.Context, key string) (int64, error) {
	attrStr := c.Request.Form.Get(key)
	attr, err := strconv.ParseInt(attrStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return attr, nil
}

func getPage(c *gin.Context) *entity.Page {
	pageNoStr := c.Query("page_no")
	pageSizeStr := c.Query("page_size")
	pageNo, err := strconv.ParseInt(pageNoStr, 10, 64)
	if err != nil {
		return nil
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		return nil
	}
	return &entity.Page{
		PageSize: int(pageSize),
		PageNo:   int(pageNo),
	}
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

	savePath := path.Join(config.Instance.WorkDir, file.Filename)
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

	userID := c.Value("userID").(uint)
	user, err := h.groupSvc.UserSvc.GetUser(c.Request.Context(), userID)
	if err != nil {
		return nil, err
	}

	info := &entity.AppInfo{
		User:               user,
		UUID:               uuid.NewString(),
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

func (h *Handler) HandleGetMe(c *gin.Context) {
	userID := c.Value("userID").(uint)
	user, err := h.groupSvc.UserSvc.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get info internal error"})
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) HandleChangePassword(c *gin.Context) {
	userID := c.Value("userID").(uint)
	err := h.groupSvc.UserSvc.ChangePassword(c.Request.Context(), userID, c.PostForm("password"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get info internal error"})
	}
	c.JSON(http.StatusOK, response.NewSucceedBaseResp(messages.ModifySucceed))
}

func (h *Handler) HandleRunApp(c *gin.Context) {
	appInfo, err := h.validateParams(c)
	if appInfo == nil || err != nil {
		return
	}

	result, err := h.groupSvc.TaskSvc.SubmitApp(c.Request.Context(), appInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "submit app internal error"})
		return
	}

	switch result {
	case entity.SubmitAppResultSystemBusy:
		c.JSON(http.StatusOK, &response.SubmitAppResponse{
			BaseResp: response.NewErrorBaseResp(messages.SystemBusy, messages.ErrCodeCommon),
		})
	case entity.SubmitAppResultSucceed:
		c.JSON(http.StatusOK, &response.SubmitAppResponse{
			BaseResp: response.NewSucceedBaseResp(messages.SubmitSucceed),
		})
	case entity.SubmitAppResultSystemTaskExists:
		c.JSON(http.StatusOK, &response.SubmitAppResponse{
			BaseResp: response.NewErrorBaseResp(messages.TaskAlreadyExists, messages.ErrCodeCommon),
		})
	}
}

func (h *Handler) HandleListAppTasks(c *gin.Context) {
	page := getPage(c)
	if page == nil {
		page = &entity.Page{
			PageSize: DefaultPageSize,
			PageNo:   DefaultPageNo,
		}
	}

	userID := c.Value("userID").(uint)

	user, err := h.groupSvc.UserSvc.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	logrus.Infof("[Handler][HandleListAppTasks] request page: %s", utils.FormatJsonString(page))
	resp, err := h.groupSvc.TaskSvc.ListAppTasks(c.Request.Context(), userID, user.Role, page)
	if err != nil {
		logrus.Errorf("[Handler][HandleListAppTasks] internal error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) HandleGetLog(c *gin.Context) {
	logType := c.Query("log_type")
	uuid := c.Query("uuid")
	if (logType != constants.LogTypeStdout && logType != constants.LogTypeStderr && logType != constants.LogTypeHurlTest) ||
		uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param"})
		return
	}
	logReader, err := h.groupSvc.TaskSvc.GetLogFile(c.Request.Context(), uuid, logType)
	if err != nil {
		logrus.Errorf("[Handler][HandleGetLog] internal error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer logReader.Close()
	if _, err = io.Copy(c.Writer, logReader); err != nil {
		logrus.Errorf("[Handler][HandleGetLog] internal error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}

func (h *Handler) HandleLogin(c *gin.Context) {
	identifier := c.PostForm("identifier")
	password := c.PostForm("password")
	req := request.LoginRequest{
		Identifier: identifier,
		Password:   password,
	}
	logrus.Infof("[Handler][HandleLogin] request: %s", utils.FormatJsonString(req))
	resp, err := h.groupSvc.UserSvc.Login(c.Request.Context(), &req)
	if err != nil {
		logrus.Errorf("[Handler][HandleLogin] internal error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, resp)
}
