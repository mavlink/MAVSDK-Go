package shell

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client ShellServiceClient
}

/*
   Send a command line.

   Parameters
   ----------
   command string


*/

func (s *ServiceImpl) Send(ctx context.Context, command string) (*SendResponse, error) {

	request := &SendRequest{}
	request.Command = command
	response, err := s.Client.Send(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Receive feedback from a sent command line.

   This subscription needs to be made before a command line is sent, otherwise, no response will be sent.


*/

func (a *ServiceImpl) Receive(ctx context.Context) (<-chan string, error) {
	ch := make(chan string)
	request := &SubscribeReceiveRequest{}
	stream, err := a.Client.SubscribeReceive(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ReceiveResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetData()
		}
	}()
	return ch, nil
}
