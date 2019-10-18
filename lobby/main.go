package main

import (
	"lemna/agent/server"
	"lemna/logger"
	"lemna/msg"
	"unicode/utf8"
	"volcano/message"
	"volcano/service"
)

func onHiMsg(id uint32, msg interface{}, from msg.Stream) {
	m := msg.(*message.HiMsg)
	logger.Debug(utf8.RuneCountInString(m.Msg), "   ", m.Msg)
	m.Msg = "I'm " + lobby.Name
	_ = from.Send(id, m)
}

var lobby *service.Service

func init() {
	lobby = service.NewService(message.SERVICE_LOBBY, server.SERVERSCHEROUND)
	lobby.Proc.Reg(&message.HiMsg{}, onHiMsg)
}

func main() {
	lobby.Main()
}
