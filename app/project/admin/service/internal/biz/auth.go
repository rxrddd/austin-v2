package biz

import (
	"context"
	"github.com/golang-jwt/jwt/v4"

	"github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/conf"
)

type AuthUseCase struct {
	key               string
	administratorRepo AdministratorRepo
}

func NewAuthUseCase(conf *conf.Auth, administratorRepo AdministratorRepo) *AuthUseCase {
	return &AuthUseCase{
		key:               conf.ApiKey,
		administratorRepo: administratorRepo,
	}
}

func (receiver *AuthUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	// 获取用户
	user, err := receiver.administratorRepo.FindLoginAdministratorByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	// 验证密码
	err = receiver.administratorRepo.VerifyPassword(ctx, user.Id, req.Password)
	if err != nil {
		return nil, err
	}
	// 生成token
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"AdministratorId":       user.Id,
			"AdministratorUsername": user.Username,
		})
	signedString, _ := claims.SignedString([]byte(receiver.key))
	return &v1.LoginReply{
		Token: signedString,
	}, nil
}
