package user

import (
	"context"
	"fmt"
	"github.com/bellingham07/go-tool/codex"
	encrypt "github.com/bellingham07/go-tool/encryt"
	"github.com/bellingham07/go-tool/errorx"
	"github.com/bellingham07/go-tool/jwtc"
	"strconv"
	"time"
	"zero-chat/api/internal/model"

	"zero-chat/api/internal/svc"
	"zero-chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var (
		user  *model.User
		token string
	)

	user, err = l.svcCtx.UserModel.QueryUser(l.ctx, fmt.Sprintf("where email = %s", req.Email))
	if err != nil {
		return nil, errorx.Internal(err, "query user error").Show()
	}

	// 校验密码
	ok := encrypt.PasswordVerify(req.Password, user.Password)
	if !ok {
		err = errorx.New("pwd error", int(codex.CodeWrongPassword), codex.CodeWrongPassword.Msg()).Show()
		return
	}

	// 颁发token
	secretKey := l.svcCtx.Config.Auth.AccessSecret
	seconds := l.svcCtx.Config.Auth.AccessExpire
	iat := time.Now().Unix()
	if token, err = jwtc.GenJwtToken(secretKey, iat, seconds, "0", strconv.FormatInt(user.Id, 10)); err != nil {
		err = errorx.New("gen token error", int(codex.CodeGenTokenErr), codex.CodeGenTokenErr.Msg())
		return
	}
	resp = &types.LoginResp{
		TokenReply: types.TokenReply{
			AccessToken:  token,
			AccessExpire: iat + seconds,
		},
	}

	return
}
