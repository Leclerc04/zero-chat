package chat

import (
	"github.com/bellingham07/go-tool/httpc"
	"net/http"

	"zero-chat/chat/api/internal/logic/chat"
	"zero-chat/chat/api/internal/svc"
)

func SubHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chat.NewSubLogic(r.Context(), svcCtx)
		err := l.Sub()
		if err != nil {
			httpc.RespError(w, r, err)
		} else {
			httpc.RespSuccess(r.Context(), w, nil)
		}
	}
}
