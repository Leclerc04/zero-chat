package contact

import (
	"net/http"
	"zero-chat/api/internal/logic/contact"
	"zero-chat/api/internal/svc"
	"zero-chat/api/internal/types"

	"github.com/leclerc04/go-tool/errorx"
	"github.com/leclerc04/go-tool/httpc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpc.RespError(w, r, errorx.BadRequest("%s:%s", errorx.CodeInvalidParams.Msg(), err.Error()).Show())
			return
		}
		l := contact.NewSearchUserLogic(r.Context(), svcCtx)
		resp, err := l.SearchUser(&req)
		if err != nil {
			httpc.RespError(w, r, err)
		} else {
			httpc.RespSuccess(r.Context(), w, resp)
		}
	}
}
