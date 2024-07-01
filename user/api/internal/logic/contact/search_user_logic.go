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

type SearchUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchUserLogic) SearchUser(req *types.SearchUserReq) (resp *types.SearchUserResp, err error) {
	query := fmt.Sprintf("id = %s", req.Uid)
	user, err := l.svcCtx.UserModel.QueryUser(l.ctx, query)
	if err != nil {
		return nil, errorx.Internal(err, err.Error()).Show()
	}

	userId := l.ctx.Value("userID").(string)
	contacts, err := l.svcCtx.ContactModel.QueryUser(l.ctx, fmt.Sprintf("owner_id = %s and contact_id = %s", userId, strconv.FormatInt(user.Id, 10)))
	if err != nil {
		return nil, errorx.Internal(err, err.Error())
	}

	if contacts.Id != 0 {
		return nil, errorx.New("this user has been added", 504, "this user has been added").Show()
	}

	resp = new(types.SearchUserResp)
	resp.Id = user.Id
	resp.Nickname = user.Nickname
	resp.Email = user.Email
	return
}
