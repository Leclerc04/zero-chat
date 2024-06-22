package chat

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-chat/api/internal/svc"
)

type CaoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaoLogic {
	return &CaoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaoLogic) Cao() error {
	// todo: add your logic here and delete this line
	fmt.Println("caojinbo")
	return nil
}
