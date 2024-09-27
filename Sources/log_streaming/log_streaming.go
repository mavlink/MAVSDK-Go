package log_streaming

import (
	"context"
	"fmt"
	"io"

	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client LogStreamingServiceClient
}

/*
   Start streaming logging data.


*/

func (s *ServiceImpl) StartLogStreaming(ctx context.Context) (*StartLogStreamingResponse, error) {

	request := &StartLogStreamingRequest{}
	response, err := s.Client.StartLogStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Stop streaming logging data.


*/

func (s *ServiceImpl) StopLogStreaming(ctx context.Context) (*StopLogStreamingResponse, error) {

	request := &StopLogStreamingRequest{}
	response, err := s.Client.StopLogStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Subscribe to logging messages


*/

func (a *ServiceImpl) LogStreamingRaw(ctx context.Context) (<-chan *LogStreamingRaw, error) {
	ch := make(chan *LogStreamingRaw)
	request := &SubscribeLogStreamingRawRequest{}
	stream, err := a.Client.SubscribeLogStreamingRaw(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				m := &LogStreamingRawResponse{}
				err := stream.RecvMsg(m)
				if err == io.EOF {
					return
				}
				if err != nil {
					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
						return
					}
					fmt.Printf("Unable to receive message: %v\n", err)
					break
				}
				ch <- m.GetLoggingRaw()
			}
		}
	}()
	return ch, nil
}
