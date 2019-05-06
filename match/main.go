package main

import (
	"fmt"
	"lemna/agent"
	"lemna/agent/arpc"
	"lemna/agent/server"
	"lemna/logger"
	"volcano/message"
	"volcano/service"
)

func Handler_HiMsg(id int32, msg interface{}, from arpc.MsgServer) {
	m := msg.(*message.HiMsg)
	m.Msg = fmt.Sprintf("hi %d,I'm %s. your msg=\"%s\"", id, match.Name, m.Msg)
	logger.Debugf("<%d>%s,%d", id, m.Msg, from.ID())
	err := match.Redis.Publish(&message.HiContent{UID: id, AID: from.ID()})
	if err != nil {
		logger.Error(err)
	}
	from.Forward(id, m)
}

func Handler_ClientLogoutMsg(id int32, msg interface{}, from arpc.MsgServer) {
	logger.Info(id, " Client logout")
}

var match *service.Service

func init() {
	match = service.NewService(message.SERVICE_MATCH, server.SERVERSCHEROUND)
	match.Rpcss.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
	match.Rpcss.Msgcenter.Reg(&agent.ClientByeMsg{}, Handler_ClientLogoutMsg)
}

func main() {
	match.Main()
}
