package logic

import (
	"context"
	"log"

	"blogsypanther/service/user/api/internal/svc"
	"blogsypanther/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
log.Println("==============LOGIN=========="); log.Println(req);log.Println("==============LOGIN==========")
  resp.Gender = "F"
  return resp, nil
}
