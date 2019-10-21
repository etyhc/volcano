package main

import (
	"fmt"
	"lemna/arpc"
	"lemna/arpc/client"
	"lemna/logger"
	"lemna/msg"
	"volcano/message"
)

type RPC struct {
	name string
	crpc client.Crpc
	proc *msg.Processor
}

func onHiMsg(t uint32, msg interface{}, from msg.Stream) {
	m := msg.(*message.HiMsg)
	logger.Info(m.Msg)
}

var rpc *RPC

func init() {
	rpc = &RPC{}
	rpc.name = "我"
	rpc.crpc.Token = "token1"
	rpc.crpc.Addr = ":9999"
	rpc.proc = msg.NewProcessor(msg.ProtoHelper{})
	rpc.proc.Reg(&message.HiMsg{}, onHiMsg)
}
func (r *RPC) Send(target uint32, msg interface{}) error {
	fmsg, err := r.proc.WrapFM(target, msg)
	if err == nil {
		return r.crpc.Send(fmsg)
	}
	return err
}

func (r *RPC) Run() {
	for {
		err := r.crpc.Login()
		if err == nil {
			logger.Infof("登录%s成功，等待输入指令", r.crpc.Addr)
			for {
				var in *arpc.ForwardMsg
				in, err = r.crpc.Recv()
				if err != nil {
					break
				}
				err = r.proc.Handle(in, r)
				if err != nil {
					logger.Error(err)
				}
			}
		}
		logger.Error(err)
		logger.Errorf("Login %s error. relogin...", r.crpc.Addr)
	}
}

func (r *RPC) Input() {
	for {
		var servertype uint32
		var msg string
		//time.Sleep(time.Second)
		fmt.Scanf("%d %s\n", &servertype, &msg)
		if servertype == 0 {
			break
		}
		err := r.Send(servertype, &message.HiMsg{Msg: msg})
		if err != nil {
			logger.Error(err)
		}
	}
}

func main() {
	go rpc.Run()
	rpc.Input()
}
