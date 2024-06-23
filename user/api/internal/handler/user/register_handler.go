package user

import (
	"github.com/bellingham07/go-tool/errorx"
	"github.com/bellingham07/go-tool/httpc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"zero-chat/user/api/internal/logic/user"
	"zero-chat/user/api/internal/svc"
	"zero-chat/user/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpc.RespError(w, r, errorx.BadRequest("%s:%s", errorx.CodeInvalidParams.Msg(), err.Error()).Show())
			return
		}
		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpc.RespError(w, r, err)
		} else {
			httpc.RespSuccess(r.Context(), w, resp)
		}
	}
}
