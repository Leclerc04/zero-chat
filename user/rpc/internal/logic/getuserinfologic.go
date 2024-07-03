package logic

import (
	"context"
	"strconv"

	"zero-chat/user/rpc/internal/svc"
	"zero-chat/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	uid, _ := strconv.ParseInt(in.Uid, 10, 64)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserInfoResp{
		Id:       user.Id,
		Email:    user.Email,
		Nickname: user.Nickname,
	}, nil
}
