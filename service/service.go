package service

import (
	"flag"
	"lemna/arpc"
	"lemna/arpc/server"
	"lemna/logger"
	"lemna/msg"
	"lemna/utils"
	"volcano/message"
)

// Service  服务器通用服务封装
type Service struct {
	Name string //服务器名字
	Proc *msg.Processor
	Srpc server.Srpc

	addr *string //参数，代理地址
	h    *bool   //参数，帮助
}

// NewService 新服务器rpc服务
func NewService(sid message.SERVICE, sche uint32) *Service {
	ret := &Service{}
	ret.addr = flag.String("addr", ":10000", "代理服务器地址")
	ret.h = flag.Bool("h", false, "this help")
	ret.Name = sid.String()
	ret.Srpc.Info.Type = uint32(sid)
	ret.Srpc.Info.Sche = sche
	ret.Proc = msg.NewProcessor(msg.ProtoHelper{})
	return ret
}

func (s *Service) Send(target uint32, msg interface{}) error {
	fmsg, err := s.Proc.WrapFM(target, msg)
	if err == nil {
		return s.Srpc.Send(fmsg.(*arpc.ForwardMsg))
	}
	return err
}

// Main 运行服务
//      启动rpc服务，然后发布自己的信息
func (s *Service) Main() {
	flag.Parse()
	if *s.h {
		flag.Usage()
		return
	}
	s.Srpc.Info.Addr = utils.PublishTCPAddr(*s.addr)
	s.Srpc.Addr = *s.addr
	over := make(chan int)
	go func() {
		err := s.Srpc.Connect()
		if err == nil {
			for {
				var in *arpc.ForwardMsg
				in, err = s.Srpc.Recv()
				if err != nil {
					break
				}
				err = s.Proc.Handle(in, s)
				if err != nil {
					logger.Error(err)
				}
			}
		}
		logger.Error(err)
		over <- 1
	}()
	<-over
}
