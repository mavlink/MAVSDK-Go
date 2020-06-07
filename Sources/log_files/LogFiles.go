package log_files

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client LogFilesServiceClient
}

    func(s *ServiceImpl)GetEntries()([]*GetEntries){
        request := &GetEntriesRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetEntries(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result := response.GetLogFilesResult()
        fmt.Printf("result %v\n", result)
        if result.Result != LogFilesResult_RESULT_SUCCESS{
            fmt.Printf("Error while getting GetEntries")
        }
        

        
    }

       

    func (a *ServiceImpl) DownloadLogFile(, id, path){
    	request := &SubscribeDownloadLogFileRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeDownloadLogFile(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &DownloadLogFile{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			fmt.Printf("message %v\n", m.GetDownloadLogFile())
    		}	
    }