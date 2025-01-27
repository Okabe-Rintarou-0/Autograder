package user

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	"autograder/pkg/config"
	"autograder/pkg/dao"
	"autograder/pkg/messages"
	"autograder/pkg/model/assembler"
	"autograder/pkg/model/dbm"
	"autograder/pkg/model/entity"
	"autograder/pkg/model/request"
	"autograder/pkg/model/response"
	"autograder/pkg/utils"
)

type ServiceImpl struct {
	groupDAO *dao.GroupDAO
}

func NewService(groupDAO *dao.GroupDAO) *ServiceImpl {
	return &ServiceImpl{groupDAO}
}

func (s *ServiceImpl) ImportCanvasUsers(ctx context.Context, courseID int64) (*response.ImportCanvasUsers, error) {
	users, err := s.groupDAO.CanvasDAO.ListCourseUsers(ctx, courseID)
	if err != nil {
		logrus.Errorf("[User Service][ImportCanvasUsers] CanvasDAO.ListCourseUsers error: %v", err)
		return nil, err
	}

	logrus.Infof("[User Service][ImportCanvasUsers] CanvasDAO.ListCourseUsers: %s", utils.FormatJsonString(users))

	eg, _ := errgroup.WithContext(ctx)
	for _, user := range users {
		if user.Email == nil {
			logrus.Warnf("[User Service][ImportCanvasUsers] user's email is nil")
			continue
		}
		eg.Go(func() error {
			req := &request.RegisterRequest{
				Username: user.LoginId,
				RealName: user.ShortName,
				Email:    *user.Email,
				Password: user.LoginId,
			}
			_, err := s.Register(ctx, req)
			return err
		})
	}
	if err = eg.Wait(); err != nil {
		logrus.Errorf("[User Service][ImportCanvasUsers] register error: %v", err)
		return nil, err
	}
	return &response.ImportCanvasUsers{
		BaseResp: response.NewSucceedBaseResp(messages.ImportSucceed),
	}, nil
}

func (s *ServiceImpl) Register(ctx context.Context, request *request.RegisterRequest) (*response.RegisterResponse, error) {
	user, err := s.groupDAO.UserDAO.Find(ctx, &dbm.UserFilter{
		Username: &request.Username,
		Email:    &request.Email,
	})
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
	identifier := request.Identifier
	user, err := s.groupDAO.UserDAO.Find(ctx, &dbm.UserFilter{
		Username: &identifier,
		Email:    &identifier,
	})
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
	user, err := s.groupDAO.UserDAO.Find(ctx, &dbm.UserFilter{
		ID: &userID,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("[User Service][GetMe] call UserDAO.FindByID error %+v", err)
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return assembler.ConvertUserDbmToEntity(user), nil
}

func (s *ServiceImpl) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := s.groupDAO.UserDAO.Find(ctx, &dbm.UserFilter{
		Username: &username,
	})
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
	user, err := s.groupDAO.UserDAO.Find(ctx, &dbm.UserFilter{
		ID: &userID,
	})
	if err != nil {
		logrus.Errorf("[User Service][ChangePassword] call UserDAO.FindByID error %+v", err)
		return err
	}
	user.Password = utils.Md5(newPassword)
	return s.groupDAO.UserDAO.Save(ctx, user)
}

func (s *ServiceImpl) ListUsers(ctx context.Context, keyword string, page *entity.Page) (*response.ListUsersResponse, error) {
	var filter *dbm.UserFilter
	if len(keyword) > 0 {
		keyword += "%"
		filter = &dbm.UserFilter{
			RealName: &keyword,
			Username: &keyword,
			Email:    &keyword,
			Or:       true,
		}
	}
	modelPage, err := s.groupDAO.UserDAO.ListByPage(ctx, filter, page.ToDBM())
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
