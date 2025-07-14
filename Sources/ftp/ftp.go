package ftp

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client FtpServiceClient
}

/*
Download Downloads a file to local directory.
*/
func (a *ServiceImpl) Download(
	ctx context.Context,
	remoteFilePath string,
	localDir string,
	useBurst bool,

) (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeDownloadRequest{
		RemoteFilePath: remoteFilePath,
		LocalDir:       localDir,
		UseBurst:       useBurst,
	}
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
				log.Fatalf("Unable to receive Download messages, err: %v", err)
			}
			ch <- m.GetProgressData()
		}
	}()
	return ch, nil
}

/*
Upload Uploads local file to remote directory.
*/
func (a *ServiceImpl) Upload(
	ctx context.Context,
	localFilePath string,
	remoteDir string,

) (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeUploadRequest{
		LocalFilePath: localFilePath,
		RemoteDir:     remoteDir,
	}
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
				log.Fatalf("Unable to receive Upload messages, err: %v", err)
			}
			ch <- m.GetProgressData()
		}
	}()
	return ch, nil
}

/*
ListDirectory Lists items from a remote directory.
*/
func (s *ServiceImpl) ListDirectory(
	ctx context.Context,
	remoteDir string,

) (*ListDirectoryResponse, error) {
	request := &ListDirectoryRequest{
		RemoteDir: remoteDir,
	}
	response, err := s.Client.ListDirectory(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
CreateDirectory Creates a remote directory.
*/
func (s *ServiceImpl) CreateDirectory(
	ctx context.Context,
	remoteDir string,

) (*CreateDirectoryResponse, error) {
	request := &CreateDirectoryRequest{
		RemoteDir: remoteDir,
	}
	response, err := s.Client.CreateDirectory(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RemoveDirectory Removes a remote directory.
*/
func (s *ServiceImpl) RemoveDirectory(
	ctx context.Context,
	remoteDir string,

) (*RemoveDirectoryResponse, error) {
	request := &RemoveDirectoryRequest{
		RemoteDir: remoteDir,
	}
	response, err := s.Client.RemoveDirectory(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RemoveFile Removes a remote file.
*/
func (s *ServiceImpl) RemoveFile(
	ctx context.Context,
	remoteFilePath string,

) (*RemoveFileResponse, error) {
	request := &RemoveFileRequest{
		RemoteFilePath: remoteFilePath,
	}
	response, err := s.Client.RemoveFile(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Rename Renames a remote file or remote directory.
*/
func (s *ServiceImpl) Rename(
	ctx context.Context,
	remoteFromPath string,
	remoteToPath string,

) (*RenameResponse, error) {
	request := &RenameRequest{
		RemoteFromPath: remoteFromPath,
		RemoteToPath:   remoteToPath,
	}
	response, err := s.Client.Rename(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
AreFilesIdentical Compares a local file to a remote file using a CRC32 checksum.
*/
func (s *ServiceImpl) AreFilesIdentical(
	ctx context.Context,
	localFilePath string,
	remoteFilePath string,

) (*AreFilesIdenticalResponse, error) {
	request := &AreFilesIdenticalRequest{
		LocalFilePath:  localFilePath,
		RemoteFilePath: remoteFilePath,
	}
	response, err := s.Client.AreFilesIdentical(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetTargetCompid Set target component ID. By default it is the autopilot.
*/
func (s *ServiceImpl) SetTargetCompid(
	ctx context.Context,
	compid uint32,

) (*SetTargetCompidResponse, error) {
	request := &SetTargetCompidRequest{
		Compid: compid,
	}
	response, err := s.Client.SetTargetCompid(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
