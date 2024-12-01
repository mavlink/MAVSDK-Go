package ftp_server

import (
	"context"
)

type ServiceImpl struct {
	Client FtpServerServiceClient
}

/*
SetRootDir Set root directory.

	This is the directory that can then be accessed by a client.
	The directory needs to exist when this is called.
	The permissions are the same as the file permission for the user running the server.
	The root directory can't be changed while an FTP process is in progress.
*/
func (s *ServiceImpl) SetRootDir(
	ctx context.Context,
	path string,

) (*SetRootDirResponse, error) {
	request := &SetRootDirRequest{
		Path: path,
	}
	response, err := s.Client.SetRootDir(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
