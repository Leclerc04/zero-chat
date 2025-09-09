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

func DelContactHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelContactReq
		if err := httpx.Parse(r, &req); err != nil {
			httpc.RespError(w, r, errorx.BadRequest("%s:%s", errorx.CodeInvalidParams.Msg(), err.Error()).Show())
			return
		}
		l := contact.NewDelContactLogic(r.Context(), svcCtx)
		err := l.DelContact(&req)
		if err != nil {
			httpc.RespError(w, r, err)
		} else {
			httpc.RespSuccess(r.Context(), w, nil)
		}
	}
}
