package ftp

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client FtpServiceClient
}

/*
   Resets FTP server in case there are stale open sessions.


*/

func (s *ServiceImpl) Reset() (*ResetResponse, error) {

	request := &ResetRequest{}
	ctx := context.Background()
	response, err := s.Client.Reset(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Downloads a file to local directory.

   Parameters
   ----------
   remoteFilePath string, localDir string
*/

func (a *ServiceImpl) Download(remoteFilePath string, localDir string) (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeDownloadRequest{}
	request.RemoteFilePath = remoteFilePath
	request.LocalDir = localDir
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
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

func (a *ServiceImpl) Upload(localFilePath string, remoteDir string) (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeUploadRequest{}
	request.LocalFilePath = localFilePath
	request.RemoteDir = remoteDir
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
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
   True
   Paths : []*string
        The found directory contents.


*/

func (s *ServiceImpl) ListDirectory(remoteDir string) (*ListDirectoryResponse, error) {
	request := &ListDirectoryRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) CreateDirectory(remoteDir string) (*CreateDirectoryResponse, error) {

	request := &CreateDirectoryRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) RemoveDirectory(remoteDir string) (*RemoveDirectoryResponse, error) {

	request := &RemoveDirectoryRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) RemoveFile(remoteFilePath string) (*RemoveFileResponse, error) {

	request := &RemoveFileRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) Rename(remoteFromPath string, remoteToPath string) (*RenameResponse, error) {

	request := &RenameRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) AreFilesIdentical(localFilePath string, remoteFilePath string) (*AreFilesIdenticalResponse, error) {
	request := &AreFilesIdenticalRequest{}
	ctx := context.Background()
	request.LocalFilePath = localFilePath
	request.RemoteFilePath = remoteFilePath
	response, err := s.Client.AreFilesIdentical(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Set root directory for MAVLink FTP server.

   Parameters
   ----------
   rootDir string


*/

func (s *ServiceImpl) SetRootDirectory(rootDir string) (*SetRootDirectoryResponse, error) {

	request := &SetRootDirectoryRequest{}
	ctx := context.Background()
	request.RootDir = rootDir
	response, err := s.Client.SetRootDirectory(ctx, request)
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

func (s *ServiceImpl) SetTargetCompid(compid uint32) (*SetTargetCompidResponse, error) {

	request := &SetTargetCompidRequest{}
	ctx := context.Background()
	request.Compid = compid
	response, err := s.Client.SetTargetCompid(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Get our own component ID.



   Returns
   -------
   False
   Compid : uint32
        Our component ID.


*/

func (s *ServiceImpl) GetOurCompid() (*GetOurCompidResponse, error) {
	request := &GetOurCompidRequest{}
	ctx := context.Background()
	response, err := s.Client.GetOurCompid(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}
