package contact

import (
	"context"
	"strconv"
	"zero-chat/api/internal/model"
	"zero-chat/api/internal/svc"
	"zero-chat/api/internal/types"

	"github.com/leclerc04/go-tool/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddContactLogic {
	return &AddContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddContactLogic) AddContact(req *types.AddContactReq) error {
	userId := l.ctx.Value("userID").(string)
	uid, _ := strconv.ParseInt(userId, 10, 64)
	contactId, _ := strconv.ParseInt(req.Uid, 10, 64)
	newContact := &model.Contacts{
		OwnerId:   uid,
		ContactId: contactId,
	}
	if err := l.svcCtx.ContactModel.Insert(l.ctx, l.svcCtx.DB, newContact); err != nil {
		return errorx.Internal(err, "database error").Show()
	}
	newContactTwo := &model.Contacts{
		OwnerId:   contactId,
		ContactId: uid,
	}
	if err := l.svcCtx.ContactModel.Insert(l.ctx, l.svcCtx.DB, newContactTwo); err != nil {
		return errorx.Internal(err, "database error").Show()
	}
	return nil
}
