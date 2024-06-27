package contact

import (
	"context"
	"fmt"
	"github.com/bellingham07/go-tool/errorx"
	"strconv"

	"zero-chat/user/api/internal/svc"
	"zero-chat/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContactLogic {
	return &GetContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContactLogic) GetContact() (resp *types.GetContactResp, err error) {
	uid := l.ctx.Value("userID").(string)
	query := fmt.Sprintf("owner_id = %s", uid)
	contacts, err := l.svcCtx.ContactModel.QueryUser(l.ctx, query)
	if err != nil {
		return nil, errorx.Internal(err, "database error").Show()

	}

	resp = new(types.GetContactResp)
	resp.List = make([]*types.User, 0, len(contacts))

	for _, contact := range contacts {
		user, err := l.svcCtx.UserModel.QueryUser(l.ctx, fmt.Sprintf("id = %s", strconv.FormatInt(contact.ContactId, 10)))
		if err != nil {
			return nil, errorx.Internal(err, "database error").Show()

		}
		tmp := &types.User{
			Id:       contact.ContactId,
			Email:    user.Email,
			Nickname: user.Nickname,
		}
		resp.List = append(resp.List, tmp)
	}

	return
}
