package log_files

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client LogFilesServiceClient
}

/*
GetEntries Get List of log files.
*/
func (s *ServiceImpl) GetEntries(
	ctx context.Context,

) (*GetEntriesResponse, error) {
	request := &GetEntriesRequest{}
	response, err := s.Client.GetEntries(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
DownloadLogFile Download log file.
*/
func (a *ServiceImpl) DownloadLogFile(
	ctx context.Context,
	entry *Entry,

	path string,

) (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeDownloadLogFileRequest{
		Entry: entry,

		Path: path,
	}
	stream, err := a.Client.SubscribeDownloadLogFile(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &DownloadLogFileResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive DownloadLogFile messages, err: %v", err)
			}
			ch <- m.GetProgress()
		}
	}()
	return ch, nil
}

/*
EraseAllLogFiles Erase all log files.
*/
func (s *ServiceImpl) EraseAllLogFiles(
	ctx context.Context,

) (*EraseAllLogFilesResponse, error) {
	request := &EraseAllLogFilesRequest{}
	response, err := s.Client.EraseAllLogFiles(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
