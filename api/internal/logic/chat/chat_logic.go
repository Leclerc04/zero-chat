package chat

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"zero-chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *ChatLogic) Chat() error {
	upgrade := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrade.Upgrade(l.w, l.r, nil)
	if err != nil {
		fmt.Println("conn websocket failed:", err)
	}

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("received err:", err)
			}
			fmt.Println("received msg:", string(msg))
			//conn.WriteMessage(websocket.TextMessage, msg)
			if err = l.svcCtx.Redis.Publish(context.Background(), "2", msg).Err(); err != nil {
				fmt.Println("publish err:", err)
				return
			}

			fmt.Println("send to redis msg:", string(msg))
		}
	}()

	return nil
}
