package handler

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"autograder/pkg/config"
	"autograder/pkg/messages"
	"autograder/pkg/model/constants"
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

func getPage(c *gin.Context) *entity.Page {
	page := entity.Page{}
	if err := c.ShouldBind(&page); err != nil {
		page.PageSize = DefaultPageSize
		page.PageNo = DefaultPageNo
	}
	return &page
}

func (h *Handler) validateParams(c *gin.Context) (*entity.AppInfo, error) {
	req := request.SubmitAppRequest{}
	if err := c.Bind(&req); err != nil {
		logrus.Errorf("[validateParams] failed to bind request: %v", err)
		return nil, err
	}

	file := req.File
	fileExt := filepath.Ext(file.Filename)
	if fileExt != ".zip" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type, only zip files are allowed"})
		return nil, fmt.Errorf("invalid file type, only zip files are allowed")
	}

	savePath := path.Join(config.Instance.WorkDir, file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
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
		AuthenticationType: entity.AuthenticationType(req.AuthenticationType),
		JDKVersion:         req.JdkVersion,
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

func (h *Handler) HandleGetCourses(c *gin.Context) {
	courses, err := h.groupSvc.CanvasSvc.ListCourses(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get info internal error"})
	}
	c.JSON(http.StatusOK, courses)
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

func (h *Handler) HandleListUsers(c *gin.Context) {
	page := getPage(c)

	logrus.Infof("[Handler][HandleListUsers] request page: %s", utils.FormatJsonString(page))
	resp, err := h.groupSvc.UserSvc.ListUsers(c.Request.Context(), page)
	if err != nil {
		logrus.Errorf("[Handler][HandleListUsers] internal error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) HandleGetLog(c *gin.Context) {
	req := request.GetLogRequest{}
	if err := c.Bind(&req); err != nil {
		logrus.Errorf("[Handler][HandleGetLog] request bind error: %+v", err)
		return
	}
	logType := req.LogType
	UUID := req.UUID
	if (logType != constants.LogTypeStdout && logType != constants.LogTypeStderr && logType != constants.LogTypeHurlTest) ||
		UUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param"})
		return
	}
	logReader, err := h.groupSvc.TaskSvc.GetLogFile(c.Request.Context(), UUID, logType)
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
	req := request.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		logrus.Errorf("[Handler][HandleRegister] request bind error: %+v", err)
		return
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

func (h *Handler) HandleRegister(c *gin.Context) {
	req := request.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		logrus.Errorf("[Handler][HandleRegister] request bind error: %+v", err)
		return
	}
	logrus.Infof("[Handler][HandleRegister] request: %s", utils.FormatJsonString(req))
	resp, err := h.groupSvc.UserSvc.Register(c.Request.Context(), &req)
	if err != nil {
		logrus.Errorf("[Handler][HandleRegister] internal error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, resp)
}
