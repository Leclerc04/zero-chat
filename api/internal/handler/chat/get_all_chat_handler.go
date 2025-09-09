package chat

import (
	"net/http"
	"zero-chat/api/internal/logic/chat"
	"zero-chat/api/internal/svc"

	"github.com/leclerc04/go-tool/httpc"
)

func GetAllChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chat.NewGetAllChatLogic(r.Context(), svcCtx)
		resp, err := l.GetAllChat()
		if err != nil {
			httpc.RespError(w, r, err)
		} else {
			httpc.RespSuccess(r.Context(), w, resp)
		}
	}
}
