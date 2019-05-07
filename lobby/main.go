package main

import (
	"lemna/agent"
	"lemna/agent/arpc"
	"lemna/agent/server"
	"lemna/logger"
	"unicode/utf8"
	"volcano/message"
	"volcano/service"
)

func onHiMsg(id int32, msg interface{}, from arpc.MsgServer) {
	m := msg.(*message.HiMsg)
	logger.Debug(utf8.RuneCountInString(m.Msg), "   ", m.Msg)
	m.Msg = "I'm " + lobby.Name
	from.Forward(id, m)
}

func onInvalidTargetMsg(fromid int32, msg interface{}, from arpc.MsgServer) {
	logger.Info(fromid, " logout")
}

var lobby *service.Service

func init() {
	lobby = service.NewService(message.SERVICE_LOBBY, server.SERVERSCHEROUND)
	lobby.Rpcss.Msgcenter.Reg(&message.HiMsg{}, onHiMsg)
	lobby.Rpcss.Msgcenter.Reg(&agent.InvalidTargetMsg{}, onInvalidTargetMsg)
}

func main() {
	lobby.Main()
}
