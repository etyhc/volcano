package main

import (
	"context"
	"fmt"
	"lemna/agent/arpc"
	"lemna/logger"
	"time"
	"volcano/message"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	name      string
	addr      string
	stream    arpc.Client_ForwardClient
	msgcenter *arpc.MsgCenter
}

func (c *Client) Broadcast(targets []int32, msg interface{}) error {
	return fmt.Errorf("unsupport")
}

func (c *Client) Forward(target int32, msg interface{}) error {
	send, err := client.msgcenter.WrapFM(target, msg.(proto.Message))
	if err == nil {
		err = c.stream.Send(send)
	}
	return err
}

func (c *Client) GetPeerAddr() (string, bool) {
	return "", false
}

func Handler_HiMsg(t int32, msg interface{}, from arpc.MsgServer) {
	m := msg.(*message.HiMsg)
	logger.Info(m.Msg)
}

var client *Client

func init() {
	client = &Client{name: "我", addr: ":9999", msgcenter: arpc.NewMsgCenter()}
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
		ac := arpc.NewClientClient(conn)
		ctx := context.Background()
		var header metadata.MD
		_, err = ac.Login(ctx, &arpc.LoginMsg{Token: "token1"}, grpc.Header(&header))
		if err == nil {
			client.stream, err = ac.Forward(metadata.NewOutgoingContext(ctx, header))
			if err == nil {
				logger.Info("agent is alive.")
				for {
					var in *arpc.ForwardMsg
					in, err = client.stream.Recv()
					if err != nil {
						break
					}
					err = client.msgcenter.Handle(in, client)
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
		err := client.Forward(servertype, &message.HiMsg{Msg: msg})
		if err != nil {
			logger.Error(err)
		}
	}
}

func main() {
	go client.Run()
	client.Input()
}
