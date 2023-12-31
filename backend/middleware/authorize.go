package middleware

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/tokenprovider/jwt"
	userstorage "backend/module/user/userstore"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Requester interface {
	GetUserId() string
	GetEmail() string
	GetRoleId() string
	IsHasFeature(featureCode string) bool
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//Authorization : Bearer {token}
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB
func RequireAuth(appCtx appctx.AppContext) func(ctx *gin.Context) {

	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()

		store := userstorage.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(
			c.Request.Context(),
			map[string]interface{}{
				"id": payload.UserId,
			},
			"Role.RoleFeatures",
		)

		if err != nil {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			panic(err)
		}

		if !user.IsActive {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		c.Set(common.CurrentUserStr, user)
		c.Next()
	}

}
