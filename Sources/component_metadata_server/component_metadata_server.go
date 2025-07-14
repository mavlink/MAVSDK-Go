package component_metadata_server

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client ComponentMetadataServerServiceClient
}
    /*
         Provide metadata (can only be called once)

         Parameters
         ----------
         metadata []*Metadata

         
    */

    func(s *ServiceImpl)SetMetadata(ctx context.Context, metadata []*Metadata)(*SetMetadataResponse, error){
        
        request := &SetMetadataRequest{}
    	request.Metadata = metadata
            
        response, err := s.Client.SetMetadata(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       