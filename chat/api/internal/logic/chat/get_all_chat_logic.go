package chat

import (
	"context"

	"zero-chat/chat/api/internal/svc"
	"zero-chat/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllChatLogic {
	return &GetAllChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllChatLogic) GetAllChat() (resp *types.GetAllChatResp, err error) {

	return
}
