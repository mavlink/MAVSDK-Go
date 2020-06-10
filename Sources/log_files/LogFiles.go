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

func (s *ServiceImpl) GetEntries() []*Entry {
	request := &GetEntriesRequest{}
	ctx := context.Background()
	response, err := s.Client.GetEntries(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}

	result := response.GetLogFilesResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != LogFilesResult_RESULT_SUCCESS {
		fmt.Printf("Error while getting GetEntries")
	}

	return response.GetEntries()

}

/*
   Download log file.

   Parameters
   ----------
   id uint32, path string
*/

func (a *ServiceImpl) DownloadLogFile(id uint32, path string) <-chan *ProgressData {
	ch := make(chan *ProgressData)
	request := &SubscribeDownloadLogFileRequest{}
	request.Id = id
	request.Path = path
	ctx := context.Background()
	stream, err := a.Client.SubscribeDownloadLogFile(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
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
			fmt.Printf("message %v\n", m)
		}
	}()
	return ch
}
