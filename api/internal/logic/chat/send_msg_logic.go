package chat

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
	"time"
	"zero-chat/api/internal/svc"
)

type SendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
	w      http.ResponseWriter
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request, w http.ResponseWriter) *SendMsgLogic {
	return &SendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
		w:      w,
	}
}

// 防止伪跨域请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (l *SendMsgLogic) SendMsg() (err error) {
	//var (
	//	wg = sync.WaitGroup{}
	//)
	//
	ws, err := upGrade.Upgrade(l.w, l.r, nil)
	if err != nil {
		fmt.Println("update err:", err)
	}
	//defer func(ws *websocket.Conn) {
	//	err = ws.Close()
	//	if err != nil {
	//		fmt.Println("close error:", err)
	//		return
	//	}
	//}(ws)
	//MsgHandler(ws, l.svcCtx.Redis, l.ctx, &wg)
	//return
	for {
		fmt.Println("start pub")
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read msg err:", err)
		}
		if err = l.svcCtx.Redis.Publish(l.ctx, "2", msg).Err(); err != nil {
			fmt.Println("error")
		}
	}

	//return
}

func MsgHandler(ws *websocket.Conn, rds *redis.Client, ctx context.Context, wg *sync.WaitGroup) {

	fmt.Println("msg handler")
	// todo 把这里订阅redis消息另起协程，后面的写入操作，使用wg.wait(),等待
	msg, err := Subscribe(ctx, rds, "websocket")
	if err != nil {
		fmt.Println("msg handler error", err)
		return
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	if err = ws.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
		fmt.Println("write msg error:", err)
	}
	fmt.Println("message send")

}

func Publish(ctx context.Context, rds *redis.Client, channel string, msg string) error {
	fmt.Println("publish...", msg)
	if err := rds.Publish(ctx, channel, msg).Err(); err != nil {
		fmt.Println("publish err:", err)
	}
	return nil
}

// Subscribe 订阅redis消息
func Subscribe(ctx context.Context, rds *redis.Client, channel string) (string, error) {
	sub := rds.Subscribe(ctx, channel)
	fmt.Println(2)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println(3)
	if err != nil {
		fmt.Println("receive error:", err)
		return "", err
	}
	fmt.Println("subscribe...", msg)
	return msg.Payload, err
}
