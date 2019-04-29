package main

import (
	"fmt"
	"lemna/agent"
	"lemna/logger"
	"unicode/utf8"
	"volcano/message"
	"volcano/service"
)

func Handler_HiMsg(id int32, msg interface{}) {
	m := msg.(*message.HiMsg)
	logger.Info(utf8.RuneCountInString(m.Msg), "   ", m.Msg)
	m.Msg = fmt.Sprintf("hi %d,I'm %s.", id, match.Name)
	match.Rpcss.Send(id, m)
}

func Handler_ClientLogoutMsg(id int32, msg interface{}) {
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
