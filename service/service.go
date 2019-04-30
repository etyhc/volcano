package service

import (
	"flag"
	"fmt"
	"lemna/agent"
	"lemna/agent/rpc"
	"lemna/content/redis"
	contentrpc "lemna/content/rpc"
	"lemna/logger"
	"volcano/message"
)

type Service struct {
	Rpcss   *rpc.ServerService
	Redis   *redis.Channel
	Name    string
	info    agent.ServerInfo
	addr    *string
	channel *string
	h       *bool
}

func NewService(sid message.SERVICE, sche int32) *Service {
	ret := &Service{}
	ret.addr = flag.String("addr", ":1000"+fmt.Sprint(int32(sid)), "要绑定的地址")
	ret.channel = flag.String("chan", contentrpc.SERVERADDR, "发布自己的内容服务器地址")
	ret.h = flag.Bool("h", false, "this help")
	ret.Name = sid.String()
	ret.Rpcss = &rpc.ServerService{
		Addr:      *ret.addr,
		Typeid:    int32(sid),
		Msgcenter: rpc.NewMsgCenter()}
	ret.info.Type = ret.Rpcss.Typeid
	ret.info.Sche = sche
	ret.Redis = &redis.Channel{Addr: redis.REDISADDR}
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
	channel := &contentrpc.Channel{Addr: *s.channel}
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
