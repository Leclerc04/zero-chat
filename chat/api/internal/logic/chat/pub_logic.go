package chat

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-chat/chat/api/internal/svc"
)

type PubLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PubLogic {
	return &PubLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PubLogic) Pub() error {
	l.svcCtx.Redis.Publish(l.ctx, "cao", "123")
	return nil
}
