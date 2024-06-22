package chat

import (
	"github.com/bellingham07/go-tool/httpc"
	"net/http"

	"zero-chat/api/internal/logic/chat"
	"zero-chat/api/internal/svc"
)

func CaoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chat.NewCaoLogic(r.Context(), svcCtx)
		err := l.Cao()
		if err != nil {
			httpc.RespError(w, r, err)
		} else {
			httpc.RespSuccess(r.Context(), w, nil)
		}
	}
}
