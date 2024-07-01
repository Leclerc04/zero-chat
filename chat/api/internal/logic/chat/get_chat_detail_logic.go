package chat

import (
	"context"

	"zero-chat/chat/api/internal/svc"
	"zero-chat/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatDetailLogic {
	return &GetChatDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatDetailLogic) GetChatDetail(req *types.GetChatHistoryReq) (resp *types.GetChatHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
