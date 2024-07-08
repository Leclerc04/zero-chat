package chat

import (
	"fmt"
	"net/http"
	"zero-chat/api/internal/logic/chat"
	"zero-chat/api/internal/svc"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chat.NewChatLogic(r.Context(), svcCtx, w, r)
		err := l.Chat()
		if err != nil {
			//httpc.RespError(w, r, err)
			fmt.Println("chat handler err:", err)
		}
		//} else {
		//	httpc.RespSuccess(r.Context(), w, nil)
		//}
	}
}
