package chat

import (
	"net/http"
	"zero-chat/api/internal/logic/chat"
	"zero-chat/api/internal/svc"
	"zero-chat/api/internal/types"

	"github.com/leclerc04/go-tool/errorx"
	"github.com/leclerc04/go-tool/httpc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendMsgReq
		if err := httpx.Parse(r, &req); err != nil {
			httpc.RespError(w, r, errorx.BadRequest("%s:%s", errorx.CodeInvalidParams.Msg(), err.Error()).Show())
			return
		}
		l := chat.NewSendMsgLogic(r.Context(), svcCtx)
		err := l.SendMsg(&req)
		if err != nil {
			httpc.RespError(w, r, err)
		} else {
			httpc.RespSuccess(r.Context(), w, nil)
		}
	}
}
