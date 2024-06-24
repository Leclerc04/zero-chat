package chat

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-chat/chat/api/internal/svc"
)

type SubLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubLogic {
	return &SubLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubLogic) Sub() error {
	for {
		sub := l.svcCtx.Redis.Subscribe(l.ctx, "cao")
		message, err := sub.ReceiveMessage(l.ctx)
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println(message.String())
	}
	return nil
}
