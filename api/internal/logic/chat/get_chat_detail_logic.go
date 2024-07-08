package chat

import (
	"context"
	"github.com/bellingham07/go-tool/errorx"
	"strconv"
	"zero-chat/api/internal/svc"
	"zero-chat/api/internal/types"

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
	fromId, _ := strconv.ParseInt(l.ctx.Value("userID").(string), 10, 64)
	messages, err := l.svcCtx.MessageModel.GetDetailHistory(l.ctx, req.ToUid, fromId)
	if err != nil {
		return nil, errorx.Internal(err, err.Error())
	}

	resp = new(types.GetChatHistoryResp)
	resp.List = make([]*types.Message, 0, len(messages))

	for _, message := range messages {
		tmp := &types.Message{
			Msg:    message.Msg,
			T:      message.CreatedAt.Unix(),
			SendId: message.SendId,
		}
		resp.List = append(resp.List, tmp)
	}

	return
}
