package log_files

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client LogFilesServiceClient
}

/*
   Get List of log files.



   Returns
   -------
   True
   Entries : []*Entry
        List of entries


*/

func (s *ServiceImpl) GetEntries() (*GetEntriesResponse, error) {
	request := &GetEntriesRequest{}
	ctx := context.Background()
	response, err := s.Client.GetEntries(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Download log file.

   Parameters
   ----------
   entry *Entry , path string
*/

func (a *ServiceImpl) DownloadLogFile(entry *Entry, path string) (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeDownloadLogFileRequest{}
	request.Entry = entry

	request.Path = path
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetProgress()
		}
	}()
	return ch, nil
}
