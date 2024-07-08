package contact

import (
	"context"
	"github.com/bellingham07/go-tool/errorx"
	"strconv"
	"zero-chat/api/internal/svc"
	"zero-chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelContactLogic {
	return &DelContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelContactLogic) DelContact(req *types.DelContactReq) error {
	uid, _ := strconv.ParseInt(req.Uid, 10, 64)
	if err := l.svcCtx.ContactModel.Delete(l.ctx, l.svcCtx.DB, uid); err != nil {
		return errorx.Internal(err, "database error").Show()

	}
	return nil
}
