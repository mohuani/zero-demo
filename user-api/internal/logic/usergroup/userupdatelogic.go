package usergroup

import (
	"context"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateReq) (resp *types.UserInfoResp, err error) {
	m := map[int64]string{
		1: "张三",
		2: "李四",
	}

	nickname := "unknown"

	if name, ok := m[req.UserId]; ok {
		nickname = name
	}

	return &types.UserInfoResp{
		Flag: nickname,
	}, nil
}
