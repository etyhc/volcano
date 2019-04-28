package service

import (
	"flag"
	"fmt"
	"lemna/agent/rpc"
	"lemna/config"
	configrpc "lemna/config/rpc"
	"lemna/logger"
	"volcano/message"
)

type Service struct {
	Rpcss   *rpc.ServerService
	Name    string
	info    config.ServerInfo
	addr    *string
	channel *string
	h       *bool
}

func NewService(sid message.SERVICE, sche int32) *Service {
	ret := &Service{}
	ret.addr = flag.String("addr", ":1000"+fmt.Sprint(int32(sid)), "要绑定的地址")
	ret.addr = flag.String("chan", configrpc.ConfigServerAddr, "发布自己的地址")
	ret.h = flag.Bool("h", false, "this help")
	ret.Name = sid.String()
	ret.Rpcss = &rpc.ServerService{
		Addr:      *ret.addr,
		Typeid:    int32(sid),
		Msgcenter: rpc.NewMsgCenter()}
	ret.info.Type = ret.Rpcss.Typeid
	ret.info.Sche = sche
	return ret
}

func (s *Service) Main() {
	flag.Parse()
	if *s.h {
		flag.Usage()
		return
	}
	s.Rpcss.Addr = *s.addr
	s.info.Addr = *s.addr
	channel := &configrpc.ChannelUser{Addr: *s.channel}
	err := channel.Publish(&s.info)
	if err != nil {
		logger.Error(err)
		return
	}
	err = s.Rpcss.Run()
	if err != nil {
		logger.Error(err)
	}
}
