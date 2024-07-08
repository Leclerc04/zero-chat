package contact

import (
	"github.com/bellingham07/go-tool/httpc"
	"net/http"
	"zero-chat/api/internal/logic/contact"
	"zero-chat/api/internal/svc"
)

func GetContactHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := contact.NewGetContactLogic(r.Context(), svcCtx)
		resp, err := l.GetContact()
		if err != nil {
			httpc.RespError(w, r, err)
		} else {
			httpc.RespSuccess(r.Context(), w, resp)
		}
	}
}
