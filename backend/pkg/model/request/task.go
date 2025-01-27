package request

import "mime/multipart"

type SubmitAppRequest struct {
	File               *multipart.FileHeader `form:"file" binding:"required"`
	JdkVersion         int32                 `form:"jdk_version" binding:"required"`
	AuthenticationType int64                 `form:"authentication_type" binding:"required"`
	Username           *string               `form:"username"`
}

type ListAppRunTasksRequest struct {
	PageNo   int   `form:"page_no" binding:"required"`
	PageSize int   `form:"page_size" binding:"required"`
	UserID   *uint `form:"user_id"`
}

type GetLogRequest struct {
	LogType string `form:"log_type" binding:"required"`
	UUID    string `form:"uuid" binding:"required"`
}
