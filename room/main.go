package main

import (
	"flag"
	"lemna/agent/rpc"
	"lemna/config"
	configrpc "lemna/config/rpc"
	"lemna/logger"
	"unicode/utf8"
	"volcano/message"
)

type RoomService struct {
	service *rpc.ServerService
	name    string
	info    config.ServerConfig
}

func Handler_HiMsg(id int32, msg interface{}, stream interface{}) {
	m := msg.(*message.HiMsg)
	s := stream.(rpc.Server_ForwardServer)
	logger.Info(utf8.RuneCountInString(m.Msg), "   ", m.Msg)
	m.Msg = "I'm " + room.name
	sendmsg, err := room.service.Msgcenter.WrapBroadcast([]int32{id}, m)
	if err == nil {
		err = s.Send(sendmsg)
		logger.Error(err)
	}
}

var room *RoomService
var addr *string
var h *bool

func init() {
	addr = flag.String("addr", ":10001", "要绑定的地址")
	h = flag.Bool("h", false, "this help")
	room = &RoomService{}
	room.name = "房间"
	room.service = &rpc.ServerService{Addr: *addr, Typeid: 1, Msgcenter: rpc.NewMsgCenter()}
	room.info = config.ServerConfig{Addr: room.service.Addr, Type: room.service.Typeid}
	room.service.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
}

func main() {
	flag.Parse()
	if *h {
		flag.Usage()
		return
	}
	room.service.Addr = *addr
	room.info.Addr = *addr
	finder := &configrpc.ChannelUser{Addr: configrpc.ConfigServerAddr}
	err := finder.Publish(&room.info)
	if err != nil {
		logger.Error(err)
		return
	}
	err = room.service.Run()
	if err != nil {
		logger.Error(err)
	}
}
