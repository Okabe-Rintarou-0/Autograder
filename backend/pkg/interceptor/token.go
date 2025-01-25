package interceptor

import (
	"net/http"
	"time"

	"autograder/pkg/dao"

	"autograder/pkg/config"
	"autograder/pkg/messages"
	"autograder/pkg/model/response"
	"autograder/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewRoleInterceptor(role int32, groupDAO *dao.GroupDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Value("userID").(uint)
		user, err := groupDAO.UserDAO.FindById(c.Request.Context(), userID)
		if err != nil {
			logrus.Errorf("[TokenIntercept] error %+v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		if user.Role != role {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

func NewTokenInterceptor(cfg *config.TokenConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		logrus.Infof("[TokenInterceptor] token: %s", token)
		userID, expireAt, err := utils.ParseToken(cfg.Secret, token)
		if err != nil {
			logrus.Errorf("[TokenIntercept] error %+v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		if !expireAt.After(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				response.NewErrorBaseResp(messages.LoginExpired, messages.ErrCodeCommon))
			return
		}

		c.Set("userID", userID)

		newToken, _ := utils.GenerateToken(cfg.Secret, cfg.ExpireAfter, userID)
		c.Header("Authorization", newToken)

		c.Next()
	}
}
