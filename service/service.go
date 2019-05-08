package service

import (
	"flag"
	"fmt"
	"lemna/agent/arpc"
	"lemna/agent/server"
	"lemna/content/crpc"
	"lemna/content/redis"
	"lemna/logger"
	"lemna/utils"
	"time"
	"volcano/message"
)

// Service  服务器通用服务封装
type Service struct {
	Rpcss   *arpc.ServerService //服务器rpc服务
	Redis   *redis.Channel      //redis订阅频道,用于服务器间数据的订阅发布
	Name    string              //服务器名字
	info    server.Info         //服务器信息
	addr    *string             //参数，服务器侦听地址
	channel *string             //参数，发布自己信息频道地址
	h       *bool               //参数，帮助
}

// NewService 新服务器rpc服务
func NewService(sid message.SERVICE, sche uint32) *Service {
	ret := &Service{}
	ret.addr = flag.String("addr", ":1000"+fmt.Sprint(int32(sid)), "要绑定的地址")
	ret.channel = flag.String("chan", crpc.SERVERADDR, "发布自己的内容服务器地址")
	ret.h = flag.Bool("h", false, "this help")
	ret.Name = sid.String()
	ret.Rpcss = &arpc.ServerService{
		Addr:      *ret.addr,
		Typeid:    uint32(sid),
		Msgcenter: arpc.NewMsgCenter()}
	ret.info.Type = ret.Rpcss.Typeid
	ret.info.Sche = sche
	ret.Redis = &redis.Channel{Addr: redis.REDISADDR}
	return ret
}

// Main 运行服务
//      启动rpc服务，然后发布自己的信息
func (s *Service) Main() {
	flag.Parse()
	if *s.h {
		flag.Usage()
		return
	}
	s.Rpcss.Addr = *s.addr
	s.info.Addr = utils.PublishTCPAddr(*s.addr)
	channel := &crpc.Channel{Addr: *s.channel}
	//延迟发布，否则先发布再起服务有问题
	over := make(chan int)
	go func() {
		tick := time.Tick(time.Second)
		<-tick
		err := channel.Publish(&s.info)
		if err != nil {
			logger.Error(err)
			over <- 1
		} else {
			logger.Info("Publish addr=", s.info.Addr)
		}
	}()
	go func() {
		err := s.Rpcss.Run()
		if err != nil {
			logger.Error(err)
		}
		over <- 1
	}()
	<-over
}
