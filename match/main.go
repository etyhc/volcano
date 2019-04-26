package main

import (
	"flag"
	"lemna/agent/rpc"
	"lemna/config"
	configrpc "lemna/config/rpc"
	"lemna/logger"
	"log"
	"volcano/message"
)

type MatchService struct {
	service *rpc.ServerService
	name    string
	info    config.ServerInfo
}

func Handler_HiMsg(id int32, msg interface{}, stream interface{}) {
	m := msg.(*message.HiMsg)
	log.Println(m)
	s := stream.(rpc.Server_ForwardServer)
	m.Msg = "I'm " + match.name
	sendmsg, err := match.service.Msgcenter.WrapBroadcast([]int32{id}, m)
	if err == nil {
		err = s.Send(sendmsg)
		logger.Error(err)
	}
}

var match *MatchService
var addr *string
var h *bool

func init() {
	addr = flag.String("addr", ":10002", "要绑定的地址")
	h = flag.Bool("h", false, "this help")
	match = &MatchService{}
	match.name = "匹配"
	match.service = &rpc.ServerService{Addr: *addr, Typeid: 2, Msgcenter: rpc.NewMsgCenter()}
	match.info.Type = match.service.Typeid
	match.info.Sche = config.SERVERSCHEROUND
	match.service.Msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
}

func main() {
	flag.Parse()
	if *h {
		flag.Usage()
		return
	}
	match.service.Addr = *addr
	match.info.Addr = *addr
	finder := &configrpc.ChannelUser{Addr: configrpc.ConfigServerAddr}
	err := finder.Publish(&match.info)
	if err != nil {
		logger.Error(err)
		return
	}
	err = match.service.Run()
	if err != nil {
		logger.Error(err)
	}
}
