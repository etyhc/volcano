package main

import (
	"context"
	"fmt"
	"io"
	"lemna/rpc"
	"log"
	"time"
	"volcano/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	name   string
	addr   string
	stream rpc.Server_ForwardClient
}

func Handler_HiMsg(t int32, msg interface{}, stream interface{}) {
	m := msg.(*message.HiMsg)
	log.Println(m)
}

func init() {
	rpc.MsgReg(&message.HiMsg{}, Handler_HiMsg)
}

func (client *Client) Run() {
	for {
		conn, err := grpc.Dial(client.addr)
		if err != nil {
			time.Sleep(time.Second)
			log.Println("agent is dead.")
		} else {
			ac := rpc.NewClientClient(conn)
			cont := context.Background()
			var header metadata.MD
			_, err := ac.Register(cont, &rpc.ClientRegMsg{Token: "token1"}, grpc.Header(&header))
			if err == nil {
				client.stream, err = ac.Forward(metadata.NewOutgoingContext(context.Background(), header))
				if err == nil {
					log.Println("agent is alive.")
					for {
						in, err := client.stream.Recv()
						if err != nil && err != io.EOF {
							log.Println(err)
							break
						}
						rpc.ForwardMsgHandle(in, client.stream)
					}
				}
			}
			log.Println(err)
			conn.Close()
		}
	}
}

func (client *Client) Input() {
	for {
		//time.Sleep(time.Second)
		var servertype int32
		var msg string
		fmt.Scanf("%d %s\n", &servertype, &msg)
		if servertype == 0 {
			break
		}
		send, err := rpc.ForwardMsgWrap(servertype, &message.HiMsg{Msg: msg})
		if err == nil {
			client.stream.Send(send)
		}
	}
}

func main() {
	client := Client{name: "我", addr: ":9999"}
	go client.Run()
	client.Input()
}
