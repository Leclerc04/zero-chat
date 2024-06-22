package chat

import (
	"fmt"
	"net/http"

	"zero-chat/api/internal/logic/chat"
	"zero-chat/api/internal/svc"
)

func SendMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chat.NewSendMsgLogic(r.Context(), svcCtx, r, w)
		err := l.SendMsg()

		if err != nil {
			//httpc.RespError(w, r, err)
			fmt.Println("handler error:", err)
		}
		//} else {
		//	httpc.RespSuccess(r.Context(), w, nil)
		//}
	}
}
