package user

import (
	"autograder/pkg/model/dbm"
	"context"
	"errors"

	"autograder/pkg/config"
	"autograder/pkg/dao"
	"autograder/pkg/messages"
	"autograder/pkg/model/assembler"
	"autograder/pkg/model/entity"
	"autograder/pkg/model/request"
	"autograder/pkg/model/response"
	"autograder/pkg/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	groupDAO *dao.GroupDAO
}

func NewService(groupDAO *dao.GroupDAO) *ServiceImpl {
	return &ServiceImpl{groupDAO}
}

func (s *ServiceImpl) Register(ctx context.Context, request *request.RegisterRequest) (*response.RegisterResponse, error) {
	user, err := s.groupDAO.UserDAO.FindByUsernameOrEmail(ctx, request.Username, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("[User Service][Register] call UserDAO.FindByUsernameOrEmail error %+v", err)
		return nil, err
	}
	resp := &response.RegisterResponse{}
	if user != nil {
		logrus.Errorf("[User Service][Login] user(%d)'s already exists", user.ID)
		resp.BaseResp = response.NewErrorBaseResp(messages.EmailOrUsernameExists, messages.ErrCodeCommon)
		return resp, nil
	}

	user = &dbm.User{
		Username: request.Username,
		RealName: request.RealName,
		Password: utils.Md5(request.Password),
		Email:    request.Email,
		Role:     dbm.CommonUser,
	}
	err = s.groupDAO.UserDAO.Save(ctx, user)
	if err != nil {
		logrus.Errorf("[User Service][Register] call UserDAO.Save error %+v", err)
		return nil, err
	}
	resp.BaseResp = response.NewSucceedBaseResp(messages.RegisterSucceed)
	return resp, nil
}

func (s *ServiceImpl) Login(ctx context.Context, request *request.LoginRequest) (*response.LoginResponse, error) {
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

func (s *ServiceImpl) GetUser(ctx context.Context, userID uint) (*entity.User, error) {
	user, err := s.groupDAO.UserDAO.FindById(ctx, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("[User Service][GetMe] call UserDAO.FindByID error %+v", err)
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return assembler.ConvertUserDbmToEntity(user), nil
}

func (s *ServiceImpl) ChangePassword(ctx context.Context, userID uint, newPassword string) error {
	user, err := s.groupDAO.UserDAO.FindById(ctx, userID)
	if err != nil {
		logrus.Errorf("[User Service][ChangePassword] call UserDAO.FindByID error %+v", err)
		return err
	}
	user.Password = utils.Md5(newPassword)
	return s.groupDAO.UserDAO.Save(ctx, user)
}

func (s *ServiceImpl) ListUsers(ctx context.Context, page *entity.Page) (*response.ListUsersResponse, error) {
	modelPage, err := s.groupDAO.UserDAO.ListByPage(ctx, page.ToDBM())
	if err != nil {
		logrus.Errorf("[User Service][ListUsers] call UserDAO.ListByPage error %+v", err)
		return nil, err
	}
	resp := &response.ListUsersResponse{
		Total: modelPage.Total,
		Data:  utils.Map(modelPage.Items, assembler.ConvertUserDbmToResponse),
	}
	return resp, err
}
