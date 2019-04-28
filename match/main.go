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

var match *service.Service

func init() {
	match = service.NewService(message.SERVICE_MATCH, agent.SERVERSCHENIL)
	match.Rpcss.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
}

func main() {
	match.Main()
}
