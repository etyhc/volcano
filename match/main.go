package main

import (
	"lemna/logger"
	"lemna/rpc"
	"log"
	"volcano/message"
)

type MatchService struct {
	service *rpc.ServerService
	name    string
}

var match *MatchService

func Handler_HiMsg(id int32, msg interface{}, stream interface{}) {
	m := msg.(*message.HiMsg)
	log.Println(m)
	s := stream.(rpc.Server_ForwardServer)
	m.Msg = "I'm match"
	sendmsg, err := match.service.Msgcenter.Wrap(id, m)
	if err == nil {
		s.Send(sendmsg)
	}
}
func init() {
	match = &MatchService{service: &rpc.ServerService{Addr: ":10002", Typeid: 1, Msgcenter: rpc.NewMsgCenter()}, name: "Match"}
	match.service.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
}

func main() {
	err := match.service.Run()
	if err != nil {
		logger.Error(err)
	}
}
