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
	himsg := msg.(*message.HiMsg)
	logger.Info(utf8.RuneCountInString(himsg.Msg), "   ", himsg.Msg)
	himsg.Msg = fmt.Sprintf("hi %d,I'm %s.", fromid, room.service.Name)
	room.service.Rpcss.Send(fromid, himsg)
}

func Handler_ClientLogoutMsg(fromid int32, msg interface{}) {
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
			Handler_HiContent(hc.(*message.HiContent))
		}
	}()
	return nil
}

func init() {
	room.service = service.NewService(message.SERVICE_ROOM, agent.SERVERSCHENIL)
	room.service.Rpcss.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
	room.service.Rpcss.Msgcenter.Reg(&agent.ClientLogoutMsg{}, Handler_ClientLogoutMsg)
}

func Handler_HiContent(hc *message.HiContent) {
	himsg := message.HiMsg{Msg: fmt.Sprintf("hi %d,I'm %s.", hc.Uid, room.service.Name)}
	room.service.Rpcss.Send(hc.Uid, &himsg)
}

func main() {
	if room.Subscribe() != nil {
		return
	}
	room.service.Main()
}
