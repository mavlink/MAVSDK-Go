import (
	"context"
	"fmt"
	"io"
)

type Service interface{
    Result() LogFilesResult_Result

}

type ServiceImpl struct{
    Client LogFilesServiceClient
}

    func(s *ServiceImpl)get_entries(self)([]*GetEntries){
     request = &GetEntriesRequest{}
         response, err = s.Client.GetEntries(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetLogFilesResult()
        fmt.Printf("result %v\n", result)
        if result.Result != LogFilesResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading GetEntries")
        }
        

        
    }

       

    func (a *ServiceImpl) DownloadLogFile(, id, path){
    	request := &DownloadLogFileRequest{}
    		ctx := context.Background()
    		stream, err := a.LogFiles.SubscribeDownloadLogFile(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &download_log_fileResponse{}
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