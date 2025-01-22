package request

import "mime/multipart"

type SubmitAppRequest struct {
	File               *multipart.FileHeader `form:"file" binding:"required"`
	JdkVersion         int32                 `form:"jdk_version" binding:"required"`
	AuthenticationType int64                 `form:"authentication_type" binding:"required"`
}

type GetLogRequest struct {
	LogType string `form:"log_type" binding:"required"`
	UUID    string `form:"uuid" binding:"required"`
}
