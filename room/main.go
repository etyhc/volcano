package main

import (
	"lemna/logger"
	"lemna/rpc"
	"log"
	"volcano/message"
)

type RoomService struct {
	service rpc.ServerService
	name    string
}

func Handler_HiMsg(id int32, msg interface{}, stream interface{}) {
	m := msg.(*message.HiMsg)
	s := stream.(rpc.Server_ForwardServer)
	log.Println(m)
	m.Msg = "I'm room"
	sendmsg, err := rpc.ForwardMsgWrap(id, m)
	if err == nil {
		s.Send(sendmsg)
	}
}

func init() {
	rpc.MsgReg(&message.HiMsg{}, Handler_HiMsg)
}

func main() {
	room := &RoomService{service: rpc.ServerService{Addr: ":10001"}, name: "Room"}
	err := room.service.Run()
	if err != nil {
		logger.Error(err)
	}
}
