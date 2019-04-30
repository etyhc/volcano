package main

import (
	"lemna/agent"
	"lemna/logger"
	"unicode/utf8"
	"volcano/message"
	"volcano/service"
)

func Handler_HiMsg(id int32, msg interface{}) {
	m := msg.(*message.HiMsg)
	logger.Debug(utf8.RuneCountInString(m.Msg), "   ", m.Msg)
	m.Msg = "I'm " + lobby.Name
	lobby.Rpcss.Send(id, m)
}

func Handler_ClientLogoutMsg(fromid int32, msg interface{}) {
	logger.Info(fromid, " logout")
}

var lobby *service.Service

func init() {
	lobby = service.NewService(message.SERVICE_LOBBY, agent.SERVERSCHEROUND)
	lobby.Rpcss.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
	lobby.Rpcss.Msgcenter.Reg(&agent.ClientLogoutMsg{}, Handler_ClientLogoutMsg)
}

func main() {
	lobby.Main()
}