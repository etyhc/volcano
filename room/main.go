package main

import (
	"fmt"
	"lemna/agent/proto"
	"lemna/agent/server"
	"lemna/arpc/msg"
	"lemna/logger"
	"unicode/utf8"
	"volcano/message"
	"volcano/service"
)

func onHiMsg(fromid uint32, msg interface{}, from msg.Stream) {
	himsg := msg.(*message.HiMsg)
	logger.Info(utf8.RuneCountInString(himsg.Msg), "   ", himsg.Msg)
	himsg.Msg = fmt.Sprintf("hi %d,I'm %s.", fromid, room.service.Name)
	_ = from.Send(fromid, himsg)
}

func onInvalidTargetMsg(fromid uint32, msg interface{}, from msg.Stream) {
	logger.Info(fromid, " logout")
}

type Room struct {
	service *service.Service
}

var room Room

func init() {
	room.service = service.NewService(message.SERVICE_ROOM, server.SERVERSCHENIL)
	room.service.Proc.Reg(&message.HiMsg{}, onHiMsg)
	room.service.Proc.Reg(&proto.InvalidTargetMsg{}, onInvalidTargetMsg)
}

func onHiContent(hc *message.HiContent) {
	himsg := message.HiMsg{Msg: fmt.Sprintf("hi %d,I'm %s.", hc.UID, room.service.Name)}
	room.service.Send(hc.UID, himsg)
}

func main() {
	room.service.Main()
}
