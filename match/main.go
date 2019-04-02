package main

import (
	"lemna/logger"
	"lemna/rpc"
	"log"
	"volcano/message"
)

type MatchService struct {
	service rpc.ServerService
	name    string
}

func Handler_HiMsg(id int32, msg interface{}, stream interface{}) {
	m := msg.(*message.HiMsg)
	log.Println(m)
	s := stream.(rpc.Server_ForwardServer)
	m.Msg = "I'm match"
	sendmsg, err := rpc.ForwardMsgWrap(id, m)
	if err == nil {
		s.Send(sendmsg)
	}
}
func init() {
	rpc.MsgReg(&message.HiMsg{}, Handler_HiMsg)
}

func main() {
	match := &MatchService{service: rpc.ServerService{Addr: ":10002"}, name: "Match"}
	err := match.service.Run()
	if err != nil {
		logger.Error(err)
	}
}
