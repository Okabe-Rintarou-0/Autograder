package entity

import (
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"autograder/pkg/config"
	"autograder/pkg/model/dbm"
	"autograder/pkg/utils"
)

type (
	AuthenticationType int
	SubmitAppResult    int
)

const (
	ByCookies AuthenticationType = 1
	ByToken   AuthenticationType = 2

	SubmitAppResultSucceed          = 0
	SubmitAppResultSystemBusy       = 1
	SubmitAppResultSystemTaskExists = 2
)

type AppInfo struct {
	User               *User
	Operator           *User
	UUID               string
	ZipFileName        string
	ProjectDirPath     string
	AuthenticationType AuthenticationType
	JDKVersion         int32
}

func (a *AppInfo) ToDBM(status int32) *dbm.AppRunTask {
	return &dbm.AppRunTask{
		UUID:       a.UUID,
		UserID:     a.User.UserID,
		OperatorID: a.Operator.UserID,
		Status:     status,
	}
}

func (a *AppInfo) GetLogDir() *LogDir {
	userID := strconv.FormatInt(int64(a.User.UserID), 10)
	return &LogDir{
		DirPath: path.Join(config.Instance.WorkDir, "logs", userID),
		UUID:    a.UUID,
	}
}

func (a *AppInfo) GetFileName() string {
	parts := strings.Split(a.ZipFileName, ".")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func (a *AppInfo) Validate() bool {
	if a.AuthenticationType != ByCookies && a.AuthenticationType != ByToken {
		return false
	}

	if a.JDKVersion != 11 && a.JDKVersion != 17 {
		return false
	}
	return true
}

func (a *AppInfo) AppPath() string {
	path := filepath.Join(config.Instance.WorkDir, "app", a.GetFileName()+"_"+a.UUID)
	return path
}

func (a *AppInfo) ZipFilePath() string {
	path := filepath.Join(config.Instance.WorkDir, a.ZipFileName)
	return path
}

func (a *AppInfo) UseCookies() bool {
	return a.AuthenticationType == ByCookies
}

func (a *AppInfo) UseToken() bool {
	return a.AuthenticationType == ByToken
}

func (a *AppInfo) FormatString() string {
	return utils.FormatJsonString(a)
}
