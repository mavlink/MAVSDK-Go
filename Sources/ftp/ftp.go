package ftp

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client FtpServiceClient
}

     /*
         Downloads a file to local directory.

         Parameters
         ----------
         remoteFilePath string, localDir string, useBurst bool
    */

    func (a *ServiceImpl) Download(ctx context.Context, remoteFilePath string, localDir string, useBurst bool) (<-chan  *ProgressData , error){
    		ch := make(chan  *ProgressData )
    		request := &SubscribeDownloadRequest{}
    		request.RemoteFilePath = remoteFilePath
        	request.LocalDir = localDir
        	request.UseBurst = useBurst
        	stream, err := a.Client.SubscribeDownload(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &DownloadResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive Download messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetProgressData()
    			}
    		}()	
    	return ch, nil
    }

     /*
         Uploads local file to remote directory.

         Parameters
         ----------
         localFilePath string, remoteDir string
    */

    func (a *ServiceImpl) Upload(ctx context.Context, localFilePath string, remoteDir string) (<-chan  *ProgressData , error){
    		ch := make(chan  *ProgressData )
    		request := &SubscribeUploadRequest{}
    		request.LocalFilePath = localFilePath
        	request.RemoteDir = remoteDir
        	stream, err := a.Client.SubscribeUpload(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &UploadResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive Upload messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetProgressData()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Lists items from a remote directory.

         Parameters
         ----------
         remoteDir string

         Returns
         -------
         False
         Data : ListDirectoryData
              The found directories and files.

         
    */


    func(s *ServiceImpl)ListDirectory(ctx context.Context, remoteDir string) (*ListDirectoryResponse, error){
        request := &ListDirectoryRequest{}
    	request.RemoteDir = remoteDir
        response, err := s.Client.ListDirectory(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Creates a remote directory.

         Parameters
         ----------
         remoteDir string

         
    */

    func(s *ServiceImpl)CreateDirectory(ctx context.Context, remoteDir string)(*CreateDirectoryResponse, error){
        
        request := &CreateDirectoryRequest{}
    	request.RemoteDir = remoteDir
        response, err := s.Client.CreateDirectory(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Removes a remote directory.

         Parameters
         ----------
         remoteDir string

         
    */

    func(s *ServiceImpl)RemoveDirectory(ctx context.Context, remoteDir string)(*RemoveDirectoryResponse, error){
        
        request := &RemoveDirectoryRequest{}
    	request.RemoteDir = remoteDir
        response, err := s.Client.RemoveDirectory(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Removes a remote file.

         Parameters
         ----------
         remoteFilePath string

         
    */

    func(s *ServiceImpl)RemoveFile(ctx context.Context, remoteFilePath string)(*RemoveFileResponse, error){
        
        request := &RemoveFileRequest{}
    	request.RemoteFilePath = remoteFilePath
        response, err := s.Client.RemoveFile(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Renames a remote file or remote directory.

         Parameters
         ----------
         remoteFromPath string

         remoteToPath string

         
    */

    func(s *ServiceImpl)Rename(ctx context.Context, remoteFromPath string, remoteToPath string)(*RenameResponse, error){
        
        request := &RenameRequest{}
    	request.RemoteFromPath = remoteFromPath
        request.RemoteToPath = remoteToPath
        response, err := s.Client.Rename(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Compares a local file to a remote file using a CRC32 checksum.

         Parameters
         ----------
         localFilePath stringremoteFilePath string

         Returns
         -------
         False
         AreIdentical : bool
              Whether the files are identical.

         
    */


    func(s *ServiceImpl)AreFilesIdentical(ctx context.Context, localFilePath string, remoteFilePath string) (*AreFilesIdenticalResponse, error){
        request := &AreFilesIdenticalRequest{}
    	request.LocalFilePath = localFilePath
        request.RemoteFilePath = remoteFilePath
        response, err := s.Client.AreFilesIdentical(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Set target component ID. By default it is the autopilot.

         Parameters
         ----------
         compid uint32

         
    */

    func(s *ServiceImpl)SetTargetCompid(ctx context.Context, compid uint32)(*SetTargetCompidResponse, error){
        
        request := &SetTargetCompidRequest{}
    	request.Compid = compid
        response, err := s.Client.SetTargetCompid(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       