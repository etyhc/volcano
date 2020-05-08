package main

import (
	"fmt"
	"lemna/agent/proto"
	"lemna/agent/server"
	"lemna/arpc/msg"
	"lemna/logger"
	"volcano/message"
	"volcano/service"
)

func onHiMsg(id uint32, msg interface{}, from msg.Stream) {
	m := msg.(*message.HiMsg)
	m.Msg = fmt.Sprintf("hi %d,I'm %s. your msg=\"%s\"", id, match.Name, m.Msg)
	logger.Debugf("<%d>%s", id, m.Msg)
	from.Send(id, m)
}

func onInvalidTargetMsg(id uint32, msg interface{}, from msg.Stream) {
	logger.Info(id, " Client logout")
}

var match *service.Service

func init() {
	match = service.NewService(message.SERVICE_MATCH, server.SERVERSCHEROUND)
	match.Proc.Reg(&message.HiMsg{}, onHiMsg)
	match.Proc.Reg(&proto.InvalidTargetMsg{}, onInvalidTargetMsg)
}

func main() {
	match.Main()
}
