package user

import (
	"autograder/pkg/model/entity"
	"context"
	"errors"

	"autograder/pkg/config"
	"autograder/pkg/dao"
	"autograder/pkg/messages"
	"autograder/pkg/model/request"
	"autograder/pkg/model/response"
	"autograder/pkg/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type serviceImpl struct {
	groupDAO *dao.GroupDAO
}

func NewService(groupDAO *dao.GroupDAO) *serviceImpl {
	return &serviceImpl{groupDAO}
}

func (s *serviceImpl) Login(ctx context.Context, request *request.LoginRequest) (*response.LoginResponse, error) {
	user, err := s.groupDAO.UserDAO.FindByUsernameOrEmail(ctx, request.Identifier, request.Identifier)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("[User Service][Login] call UserDAO.FindByUsernameOrEmail error %+v", err)
		return nil, err
	}
	resp := &response.LoginResponse{}
	logrus.Infof("[User Service][Login] get user %s", utils.FormatJsonString(user))
	if user == nil || utils.Md5(request.Password) != user.Password {
		logrus.Warnf("[User Service][Login] user(%s)'s login failed", request.Identifier)
		resp.BaseResp = response.NewErrorBaseResp(messages.WrongPasswordOrUsername, messages.ErrCodeCommon)
		return resp, nil
	}
	logrus.Infof("[User Service][Login] user(%s)'s password matched", request.Identifier)
	resp.BaseResp = response.NewSucceedBaseResp(messages.LoginSucceed)

	tokenCfg := config.Instance.Token
	resp.Token, err = utils.GenerateToken(tokenCfg.Secret, tokenCfg.ExpireAfter, user.ID)
	if err != nil {
		logrus.Errorf("[User Service][Login] call utils.GenerateToken error %+v", err)
		return nil, err
	}
	return resp, nil
}

func (s *serviceImpl) GetUser(ctx context.Context, userID uint) (*entity.User, error) {
	user, err := s.groupDAO.UserDAO.FindById(ctx, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("[User Service][GetMe] call UserDAO.FindByID error %+v", err)
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return &entity.User{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
