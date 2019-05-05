package main

import (
	"fmt"
	"lemna/agent"
	"lemna/agent/rpc"
	"lemna/logger"
	"volcano/message"
	"volcano/service"
)

func Handler_HiMsg(id int32, msg interface{}, from rpc.MsgPeer) {
	m := msg.(*message.HiMsg)
	logger.Infof("<%d>%s", id, m.Msg)
	m.Msg = fmt.Sprintf("hi %d,I'm %s. your msg=\"%s\"", id, match.Name, m.Msg)
	err := match.Redis.Publish(&message.HiContent{Uid: id})
	if err != nil {
		logger.Error(err)
	}
	from.Forward(id, m)
}

func Handler_ClientLogoutMsg(id int32, msg interface{}, from rpc.MsgPeer) {
	logger.Info(id, " Client logout")
}

var match *service.Service

func init() {
	match = service.NewService(message.SERVICE_MATCH, agent.SERVERSCHEROUND)
	match.Rpcss.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
	match.Rpcss.Msgcenter.Reg(&agent.ClientLogoutMsg{}, Handler_ClientLogoutMsg)
}

func main() {
	match.Main()
}
