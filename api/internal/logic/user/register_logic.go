package user

import (
	"context"
	"zero-chat/api/internal/model"
	"zero-chat/api/internal/svc"
	"zero-chat/api/internal/types"

	encrypt "github.com/leclerc04/go-tool/encryt"
	"github.com/leclerc04/go-tool/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user := &model.User{
		Email:    req.Email,
		Password: encrypt.PasswordHash(req.Password),
		Nickname: req.Email,
	}

	if err = l.svcCtx.UserModel.Insert(l.ctx, l.svcCtx.DB, user); err != nil {
		return nil, errorx.Internal(err, "register failed").Show()
	}
	return
}
