package main

import (
	"fmt"
	"lemna/agent"
	"lemna/agent/arpc"
	"lemna/agent/server"
	"lemna/logger"
	"unicode/utf8"
	"volcano/message"
	"volcano/service"
)

func onHiMsg(fromid int32, msg interface{}, from arpc.MsgServer) {
	himsg := msg.(*message.HiMsg)
	logger.Info(utf8.RuneCountInString(himsg.Msg), "   ", himsg.Msg)
	himsg.Msg = fmt.Sprintf("hi %d,I'm %s.", fromid, room.service.Name)
	from.Forward(fromid, himsg)
}

func onInvalidTargetMsg(fromid int32, msg interface{}, from arpc.MsgServer) {
	logger.Info(fromid, " logout")
}

type Room struct {
	service *service.Service
}

var room Room

func (r Room) Subscribe() error {
	hichan, err := room.service.Redis.Subscribe(&message.HiContent{})
	if err != nil {
		logger.Error(err)
		return err
	}
	go func() {
		for {
			hc := <-hichan
			onHiContent(hc.(*message.HiContent))
		}
	}()
	return nil
}

func init() {
	room.service = service.NewService(message.SERVICE_ROOM, server.SERVERSCHENIL)
	room.service.Rpcss.Msgcenter.Reg(&message.HiMsg{}, onHiMsg)
	room.service.Rpcss.Msgcenter.Reg(&agent.InvalidTargetMsg{}, onInvalidTargetMsg)
}

func onHiContent(hc *message.HiContent) {
	himsg := message.HiMsg{Msg: fmt.Sprintf("hi %d,I'm %s.", hc.UID, room.service.Name)}
	s := room.service.Rpcss.Get(hc.AID)
	if s != nil {
		s.Forward(hc.UID, &himsg)
	}
}

func main() {
	if room.Subscribe() != nil {
		return
	}
	room.service.Main()
}
