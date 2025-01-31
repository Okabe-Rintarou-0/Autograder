package assembler

import (
	"autograder/pkg/model/dbm"
	"autograder/pkg/model/entity"
	"autograder/pkg/model/request"
	"autograder/pkg/model/response"
)

func ConvertUserDbmToProfile(m *dbm.User) *dbm.UserProfile {
	return &dbm.UserProfile{
		ID:       m.ID,
		Username: m.Username,
		Role:     m.Role,
		RealName: m.RealName,
		Email:    m.Email,
	}
}

func ConvertUserProfileDbmToResponse(m *dbm.UserProfile) *response.UserProfile {
	return &response.UserProfile{
		ID:       m.ID,
		Username: m.Username,
		Role:     m.Role,
		RealName: m.RealName,
		Email:    m.Email,
	}
}

func ConvertAppRunTaskDbmToResponse(m *dbm.AppRunTaskWithUser) *response.AppRunTask {
	return &response.AppRunTask{
		UUID:        m.UUID,
		Error:       m.Error,
		User:        ConvertUserProfileDbmToResponse(m.User),
		Operator:    ConvertUserProfileDbmToResponse(m.Operator),
		Status:      m.Status,
		CreatedAt:   m.CreatedAt,
		Pass:        m.Pass,
		Total:       m.Total,
		TestResults: m.TestResults,
	}
}

func ConvertUserDbmToResponse(m *dbm.User) *response.User {
	return &response.User{
		Username:  m.Username,
		ID:        m.ID,
		Email:     m.Email,
		RealName:  m.RealName,
		Role:      m.Role,
		CreatedAt: m.CreatedAt,
	}
}

func ConvertUserDbmToEntity(m *dbm.User) *entity.User {
	return &entity.User{
		UserID:             m.ID,
		Username:           m.Username,
		RealName:           m.RealName,
		Email:              m.Email,
		Role:               m.Role,
		JdkVersion:         m.JdkVersion,
		AuthenticationType: m.AuthenticationType,
	}
}

func ConvertTestcaseDbmToResponse(m *dbm.Testcase) *response.Testcase {
	return &response.Testcase{
		ID:      m.ID,
		Name:    m.Name,
		Path:    m.Path,
		Status:  m.Status,
		Content: m.Content,
	}
}

func ConvertTestcaseRequestToDBM(m *request.Testcase) *dbm.Testcase {
	return &dbm.Testcase{
		Name:    m.Name,
		Path:    m.Path,
		Status:  m.Status,
		Content: m.Content,
	}
}
