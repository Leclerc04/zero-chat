package imserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
	"sync"
	"time"
)

type (
	ImServer struct {
		kafkaBroker *kafka.Reader
		rds         *redis.Client
		clients     map[string]*websocket.Conn
		Address     string
		lock        sync.Mutex
		upgraer     *websocket.Upgrader
		ctx         context.Context
	}
	SendMsgRequest struct {
		FromUid       string `json:"fromUid"`
		FromNickname  string `json:"fromNickname"`
		ToUid         string `json:"toUid"`
		ToNickname    string `json:"toNickname"`
		Body          string `json:"body"`
		TimeStamp     int64  `json:"timeStamp"`
		RemoteAddress string `json:"remoteAddress"`
	}
	LoginRequest struct {
		Uid string `json:"uid"`
	}
	SendMsgResponse struct {
		FromToken     string `json:"fromToken"`     // 消息来自谁
		Body          string `json:"body"`          // 消息内容
		RemoteAddress string `json:"remoteAddress"` // 消息远程地址
	}
	ImServerOptions func(im *ImServer)
)

// func NewImServer(rds *redis.Client, opts ImServerOptions) (*ImServer, error) {
func NewImServer(rds *redis.Client, kafkaConn *kafka.Reader) (*ImServer, error) {
	// 初始化
	//if err := broker.Init(); err != nil {
	//	return nil, err
	//}
	//if err := broker.Connect(); err != nil {
	//	return nil, err
	//}
	imServer := &ImServer{
		ctx:         context.Background(),
		rds:         rds,
		kafkaBroker: kafkaConn,
		clients:     make(map[string]*websocket.Conn, 0), // 用户：多个websocket链接
		// 初始化websocket的读取大小和写入大小
		upgraer: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	//if opts != nil {
	//	opts(imServer)
	//}
	if imServer.Address == "" {
		imServer.Address = "0.0.0.0:7272"
	}
	return imServer, nil
}

func (l *ImServer) SendMsg(r *SendMsgRequest) (*SendMsgResponse, error) {
	// websocket发送消息
	l.lock.Lock()
	defer l.lock.Unlock()
	log.Printf("send SendMsgRequest  %+v", r)
	conn := l.clients[r.ToUid] //获取用户的websocket链接
	if conn == nil {
		return nil, errors.New("user don't login")
	}
	r.TimeStamp = time.Now().Unix()
	r.RemoteAddress = conn.RemoteAddr().String()
	bodyMsg, err := json.Marshal(r)
	if err != nil {
		return nil, errors.New("send message error")
	}
	// 向websocket发送消息
	if err = conn.WriteMessage(websocket.TextMessage, bodyMsg); err != nil {
		log.Printf("send message err %v", err)
		l.clients[r.ToUid] = nil
		//log.Println(conn.Close())
		return nil, err
	}
	log.Printf("send message succes  %v", r.Body)
	return &SendMsgResponse{}, nil
}

func (l *ImServer) SubscribeTwo() {
	for {
		fmt.Println("start subscribe")
		message, err := l.kafkaBroker.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("read kafka failed:", err)
			break
		}

		r := new(SendMsgRequest)
		if err = json.Unmarshal(message.Value, r); err != nil {
			log.Printf("[Unmarshal msg err] : %+v", err)
			return
		}
		if _, err = l.SendMsg(r); err != nil {
			log.Printf("[SendMsg err] : %+v", err)
			return
		}
		log.Printf("has Subscribe msg %+v", r.Body)
	}
}

func (l *ImServer) Subscribe() {
	for {
		fmt.Println("start subscribe")
		sub := l.rds.Subscribe(l.ctx, "ws")
		message, err := sub.ReceiveMessage(l.ctx)
		if err != nil {
			log.Printf("[rds received msg errror]:%+v", err)
			return
		}
		r := new(SendMsgRequest)
		fmt.Printf("r:s%", r)
		if err = json.Unmarshal([]byte(message.Payload), r); err != nil {
			log.Printf("[Unmarshal msg err] : %+v", err)
			return
		}
		if _, err := l.SendMsg(r); err != nil {
			log.Printf("[SendMsg err] : %+v", err)
			return
		}
		log.Printf("has Subscribe msg %+v", message.Payload)
	}
}

func (l *ImServer) Run() {
	log.Printf("websocket has listens at %s", l.Address)
	http.HandleFunc("/chat/ws", l.login)
	log.Fatal(http.ListenAndServe(l.Address, nil))
}

// 用户和websocket关联，初始化该用户的ws连接，以便后续发送消息
func (l *ImServer) login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listen login")
	// 连接websocket
	conn, err := l.upgraer.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 用户登录后会发一个消息 用户已连接websocket
	msgType, message, err := conn.ReadMessage()
	if err != nil {
		log.Printf("read login message err %+v", err)
		return
	}
	log.Printf("用户已连接websocket:s%", string(message))
	log.Printf("msgType:%s", msgType)
	// 只发送文本类型的消息
	if msgType != websocket.TextMessage {
		log.Printf("read login msgType err %+v", err)
		return
	}
	fmt.Println(string(message))
	loginMsgRequest := new(LoginRequest)

	// 从连接读到的消息中获取登录的token
	if err := json.Unmarshal(message, loginMsgRequest); err != nil {
		log.Printf("json.Unmarshal msg err %+v", err)
		return
	}
	l.clients[loginMsgRequest.Uid] = conn
	fmt.Println(l.clients)
	fmt.Println("用户已登录：", string(message))
	return
}
