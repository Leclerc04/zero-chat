package chat

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"zero-chat/chat/api/internal/common/imserver"

	"zero-chat/chat/api/internal/svc"
	"zero-chat/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMsgLogic) SendMsg(req *types.SendMsgReq) error {
	userId := l.ctx.Value("userID").(string)

	r := imserver.SendMsgRequest{
		FromUid:   userId,  // qq
		ToUid:     req.Uid, // 163
		Body:      req.Msg,
		TimeStamp: time.Now().Unix(),
	}
	rJson, err := json.Marshal(r)
	if err != nil {
		log.Printf("json marshal err:%s", err)
		return err
	}
	if err = l.svcCtx.Redis.Publish(l.ctx, "ws", rJson).Err(); err != nil {
		log.Printf("publish msg err:%s", err)
		return err
	}
	return nil
}
