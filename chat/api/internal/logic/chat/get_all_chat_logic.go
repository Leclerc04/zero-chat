package chat

import (
	"context"
	"github.com/bellingham07/go-tool/errorx"
	"strconv"
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
	uid, _ := strconv.ParseInt(l.ctx.Value("userID").(string), 10, 64)
	history, err := l.svcCtx.MessageModel.GetAllHistory(l.ctx, uid)
	if err != nil {
		return nil, errorx.Internal(err, "get history error").Show()
	}

	resp = new(types.GetAllChatResp)
	resp.List = make([]*types.MessageInfo, 0, len(history))

	for _, h := range history {
		tmp := &types.MessageInfo{
			ToUid:   h.Id,
			ToUser:  strconv.FormatInt(h.Id, 10),
			LastMsg: h.Msg,
		}
		resp.List = append(resp.List, tmp)
	}

	return
}

//func (l *GetAllChatLogic)getUid(uid string) string{
//	in := pb.GetUserInfoReq{Uid: uid}
//	userInfo, err := l.svcCtx.UserC.GetUserInfo(l.ctx, &in)
//	if err != nil {
//		log.Printf("err:%s",err)
//		return ""
//	}
//
//}
