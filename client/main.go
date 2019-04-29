package main

import (
	"context"
	"fmt"
	"lemna/agent/rpc"
	"lemna/logger"
	"time"
	"unicode/utf8"
	"volcano/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	name      string
	addr      string
	stream    rpc.Client_ForwardClient
	msgcenter *rpc.MsgCenter
}

func Handler_HiMsg(t int32, msg interface{}) {
	m := msg.(*message.HiMsg)
	logger.Info(m.Msg)
}

var client *Client

func init() {
	client = &Client{name: "我", addr: ":9999", msgcenter: rpc.NewMsgCenter()}
	client.msgcenter.Reg(&message.HiMsg{}, Handler_HiMsg)
}

func (c *Client) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"token": "token1"}, nil
}

func (c *Client) RequireTransportSecurity() bool {
	return false
}

func (client *Client) Run() {
	for {
		conn, err := grpc.Dial(client.addr,
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithPerRPCCredentials(client))
		if err != nil {
			logger.Error(err)
			return
		}
		ac := rpc.NewClientClient(conn)
		ctx := context.Background()
		var header metadata.MD
		_, err = ac.Login(ctx, &rpc.LoginMsg{Token: "token1"}, grpc.Header(&header))
		if err == nil {
			client.stream, err = ac.Forward(metadata.NewOutgoingContext(ctx, header))
			if err == nil {
				logger.Info("agent is alive.")
				for {
					var in *rpc.ForwardMsg
					in, err = client.stream.Recv()
					if err != nil {
						break
					}
					err = client.msgcenter.Handle(in)
					if err != nil {
						logger.Error(err)
					}
				}
			}
		}
		conn.Close()
		logger.Error(err)
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
		send, err := client.msgcenter.WrapFM(servertype, &message.HiMsg{Msg: msg})
		if err == nil {
			err = client.stream.Send(send)
			if err != nil {
				logger.Error(err)
			}
		}
	}
}

func main() {
	go client.Run()
	client.Input()
}
