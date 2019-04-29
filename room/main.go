package main

import (
	"fmt"
	"lemna/agent"
	"lemna/logger"
	"unicode/utf8"
	"volcano/message"
	"volcano/service"
)

func Handler_HiMsg(fromid int32, msg interface{}) {
	m := msg.(*message.HiMsg)
	logger.Info(utf8.RuneCountInString(m.Msg), "   ", m.Msg)
	m.Msg = fmt.Sprintf("hi %d,I'm %s.", fromid, room.Name)
	room.Rpcss.Send(fromid, m)
}

func Handler_ClientLogoutMsg(fromid int32, msg interface{}) {
	logger.Info(fromid, " logout")
}

var room *service.Service

func init() {
	room = service.NewService(message.SERVICE_ROOM, agent.SERVERSCHENIL)
	room.Rpcss.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
	room.Rpcss.Msgcenter.Reg(&agent.ClientLogoutMsg{}, Handler_ClientLogoutMsg)
}

func main() {
	room.Main()
}
