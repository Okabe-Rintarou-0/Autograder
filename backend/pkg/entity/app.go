package entity

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"

	"autograder/pkg/config"
	"autograder/pkg/utils"
)

type AuthenticationType int

const (
	ByCookies AuthenticationType = 1
	ByToken   AuthenticationType = 2
)

type AppInfo struct {
	ZipFileName        string
	UploadTime         time.Time
	AuthenticationType AuthenticationType
	JDKVersion         int32
}

func (a *AppInfo) GetFileName() string {
	parts := strings.Split(a.ZipFileName, ".")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func (a *AppInfo) GetStudentName() string {
	parts := strings.Split(a.GetFileName(), "_")
	if len(parts) != 2 {
		return ""
	}
	return parts[0]
}

func (a *AppInfo) GetStudentID() string {
	parts := strings.Split(a.GetFileName(), "_")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func (a *AppInfo) Validate() bool {
	parts := strings.Split(a.GetFileName(), "_")
	if len(parts) != 2 {
		return false
	}
	if a.AuthenticationType != ByCookies && a.AuthenticationType != ByToken {
		return false
	}

	if a.JDKVersion != 11 && a.JDKVersion != 17 {
		return false
	}
	return true
}

func (a *AppInfo) AppPath() string {
	path := filepath.Join(config.WorkDir, a.GetFileName())
	return path
}

func (a *AppInfo) ZipFilePath() string {
	path := filepath.Join(config.WorkDir, a.ZipFileName)
	return path
}

func (a *AppInfo) GetUUID() string {
	return fmt.Sprintf("%s_%s", a.GetStudentID(), uuid.NewString())
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
