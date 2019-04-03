package main

import (
	"context"
	"fmt"
	"lemna/logger"
	"lemna/rpc"
	"log"
	"time"
	"unicode/utf8"
	"volcano/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	name      string
	addr      string
	stream    rpc.Server_ForwardClient
	msgcenter *rpc.MsgCenter
}

func Handler_HiMsg(t int32, msg interface{}, stream interface{}) {
	m := msg.(*message.HiMsg)
	logger.Info(m.Msg)
}

var client *Client

func init() {
	client = &Client{name: "我", addr: ":9999", msgcenter: rpc.NewMsgCenter()}
	client.msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
}

func (client *Client) Run() {
	for {
		conn, err := grpc.Dial(client.addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return
		} else {
			ac := rpc.NewClientClient(conn)
			ctx := context.Background()
			var header metadata.MD
			_, err := ac.Register(ctx, &rpc.ClientRegMsg{Token: "token1"}, grpc.Header(&header))
			if err == nil {
				client.stream, err = ac.Forward(metadata.NewOutgoingContext(ctx, header))
				if err == nil {
					logger.Info("agent is alive.")
					for {
						in, err := client.stream.Recv()
						if err != nil {
							log.Println(err)
							break
						}
						client.msgcenter.Handle(in, client.stream)
					}
				}
			}
			log.Println(err)
			conn.Close()
		}
		time.Sleep(time.Second * 5)
	}
}

func (client *Client) Input() {
	for {
		var servertype int32
		var msg string
		//time.Sleep(time.Second)
		fmt.Scanf("%d %s\n", &servertype, &msg)
		if servertype == 0 {
			break
		}
		logger.Info(utf8.RuneCountInString(msg), "   ", msg)
		send, err := client.msgcenter.Wrap(servertype, &message.HiMsg{Msg: msg})
		if err == nil {
			client.stream.Send(send)
		}
	}
}

func main() {
	go client.Run()
	client.Input()
}
