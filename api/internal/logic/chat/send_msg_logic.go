package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
	"zero-chat/api/internal/common/imserver"
	"zero-chat/api/internal/svc"
	"zero-chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMsgLogic) SendMsg(req *types.SendMsgReq) error {
	userId := l.ctx.Value("userID").(string)

	r := imserver.SendMsgRequest{
		FromUid:    userId,  // qq
		ToUid:      req.Uid, // 163
		ToNickname: req.Uid,
		//ToNickname: l.getNameByUid(req.Uid),
		Body:      req.Msg,
		TimeStamp: time.Now().Unix(),
	}
	rJson, err := json.Marshal(r)
	if err != nil {
		log.Printf("json marshal err:%s", err)
		return err
	}

	for i := 0; i < 3; i++ {
		if err = l.svcCtx.KafkaWriter.WriteMessages(
			l.ctx,
			// 原子操作，要么全部成功，要么全部失败
			kafka.Message{Key: []byte("zero-chat"), Value: rJson},
			//kafka.Message{Key: []byte("1"), Value: []byte("is")},
			//kafka.Message{Key: []byte("1"), Value: []byte("a boy")},
		); err != nil {
			if errors.Is(err, kafka.LeaderNotAvailable) {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				fmt.Println("write messages to kafka failed:", err)
			}
		} else {
			fmt.Println("msg sent")
			break
		}
	}

	//r := imserver.SendMsgRequest{
	//	FromUid:    userId,  // qq
	//	ToUid:      req.Uid, // 163
	//	ToNickname: l.getNameByUid(req.Uid),
	//	Body:       req.Msg,
	//	TimeStamp:  time.Now().Unix(),
	//}
	//rJson, err := json.Marshal(r)
	//if err != nil {
	//	log.Printf("json marshal err:%s", err)
	//	return err
	//}
	//if err = l.svcCtx.Redis.Publish(l.ctx, "ws", rJson).Err(); err != nil {
	//	log.Printf("publish msg err:%s", err)
	//	return err
	//}
	// store msg to db
	//sId, _ := strconv.ParseInt(userId, 10, 64)
	//rId, _ := strconv.ParseInt(req.Uid, 10, 64)
	//message := &model.Message{
	//	SendId:    sId,
	//	ReceiveId: rId,
	//	Msg:       req.Msg,
	//}
	//if err = l.svcCtx.MessageModel.Insert(l.ctx, l.svcCtx.DB, message); err != nil {
	//	return errorx.Internal(err, "store msg error").Show()
	//}
	return nil
}

//func (l *SendMsgLogic) getNameByUid(uid string) string {
//	in := pb.GetUserInfoReq{Uid: uid}
//	userInfo, err := l.svcCtx.UserC.GetUserInfo(l.ctx, &in)
//	if err != nil {
//		log.Printf("err:%s", err)
//		return ""
//	}
//	return userInfo.Nickname
//}
