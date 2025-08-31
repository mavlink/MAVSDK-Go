package shell

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client ShellServiceClient
}

/*
Send Send a command line.
*/
func (s *ServiceImpl) Send(
	ctx context.Context,
	command string,

) (*SendResponse, error) {
	request := &SendRequest{
		Command: command,
	}
	response, err := s.Client.Send(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Receive Receive feedback from a sent command line.

	This subscription needs to be made before a command line is sent, otherwise, no response will be sent.
*/
func (a *ServiceImpl) Receive(
	ctx context.Context,

) (<-chan string, error) {
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
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive Receive messages, err: %v", err)
			}
			ch <- m.GetData()
		}
	}()
	return ch, nil
}
