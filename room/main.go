package main

import (
	"fmt"
	"lemna/config"
	"lemna/logger"
	"unicode/utf8"
	"volcano/message"
	"volcano/service"
)

func Handler_HiMsg(id int32, msg interface{}) {
	m := msg.(*message.HiMsg)
	logger.Info(utf8.RuneCountInString(m.Msg), "   ", m.Msg)
	m.Msg = fmt.Sprintf("hi %d,I'm %s.", id, room.Name)
	room.Rpcss.Send(id, m)
}

var room *service.Service

func init() {
	room = service.NewService(message.SERVICE_ROOM, config.SERVERSCHENIL)
	room.Rpcss.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
}

func main() {
	room.Main()
}
