package main

import (
	"lemna/logger"
	"lemna/rpc"
	"unicode/utf8"
	"volcano/message"
)

type RoomService struct {
	service *rpc.ServerService
	name    string
}

var room *RoomService

func Handler_HiMsg(id int32, msg interface{}, stream interface{}) {
	m := msg.(*message.HiMsg)
	s := stream.(rpc.Server_ForwardServer)
	logger.Info(utf8.RuneCountInString(m.Msg), "   ", m.Msg)
	m.Msg = "I'm 房间"
	sendmsg, err := room.service.Msgcenter.Wrap(id, m)
	if err == nil {
		s.Send(sendmsg)
	}
}

func init() {
	room = &RoomService{service: &rpc.ServerService{Addr: ":10001", Typeid: 1, Msgcenter: rpc.NewMsgCenter()}, name: "Room"}
	room.service.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
}

func main() {
	err := room.service.Run()
	if err != nil {
		logger.Error(err)
	}
}
