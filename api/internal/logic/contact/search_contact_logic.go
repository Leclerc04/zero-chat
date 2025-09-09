package contact

import (
	"context"
	"fmt"
	"strconv"
	"zero-chat/api/internal/svc"
	"zero-chat/api/internal/types"

	"github.com/leclerc04/go-tool/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchContactLogic {
	return &SearchContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchContactLogic) SearchContact(req *types.SearchContactReq) (resp *types.SearchContactResp, err error) {
	userId := l.ctx.Value("userID").(string)
	uid, _ := strconv.Atoi(userId)
	query := fmt.Sprintf("nickname like '%%%s%%'	", req.Key)
	users, err := l.svcCtx.UserModel.QueryUsersByKey(l.ctx, query, uid)
	if err != nil {
		return nil, errorx.Internal(err, "query user info failed").Show()
	}

	resp = new(types.SearchContactResp)
	resp.List = make([]*types.User, 0, len(users))

	for _, user := range users {
		tmp := &types.User{
			Id:       user.Id,
			Email:    user.Email,
			Nickname: user.Nickname,
		}
		resp.List = append(resp.List, tmp)
	}

	return
}
