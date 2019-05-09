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

func onHiMsg(id uint32, msg interface{}, from arpc.MsgStream) {
	m := msg.(*message.HiMsg)
	logger.Debug(utf8.RuneCountInString(m.Msg), "   ", m.Msg)
	m.Msg = "I'm " + lobby.Name
	_ = from.Forward(id, m)
}

func onInvalidTargetMsg(fromid uint32, msg interface{}, from arpc.MsgStream) {
	logger.Info(fromid, " logout")
}

var lobby *service.Service

func init() {
	lobby = service.NewService(message.SERVICE_LOBBY, server.SERVERSCHEROUND)
	lobby.Mc.Reg(&message.HiMsg{}, onHiMsg)
	lobby.Mc.Reg(&agent.InvalidTargetMsg{}, onInvalidTargetMsg)
}

func main() {
	lobby.Main()
}
