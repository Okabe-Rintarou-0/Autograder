package assembler

import (
	"autograder/pkg/model/dbm"
	"autograder/pkg/model/entity"
	"autograder/pkg/model/response"
)

func ConvertAppRunTaskDbmToResponse(m *dbm.AppRunTaskWithUser) *response.AppRunTask {
	return &response.AppRunTask{
		UUID:        m.UUID,
		UserID:      m.UserID,
		Username:    m.Username,
		Email:       m.Email,
		RealName:    m.RealName,
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
		UserID:   m.ID,
		Username: m.Username,
		Email:    m.Email,
		Role:     m.Role,
		RealName: m.RealName,
	}
}
