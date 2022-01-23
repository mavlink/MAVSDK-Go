package info

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client InfoServiceClient
}
    /*
         Get flight information of the system.

         

         Returns
         -------
         False
         FlightInfo : FlightInfo
              Flight information of the system

         
    */


    func(s *ServiceImpl)GetFlightInformation() (*GetFlightInformationResponse, error){
        request := &GetFlightInformationRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetFlightInformation(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Get the identification of the system.

         

         Returns
         -------
         False
         Identification : Identification
              Identification of the system

         
    */


    func(s *ServiceImpl)GetIdentification() (*GetIdentificationResponse, error){
        request := &GetIdentificationRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetIdentification(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Get product information of the system.

         

         Returns
         -------
         False
         Product : Product
              Product information of the system

         
    */


    func(s *ServiceImpl)GetProduct() (*GetProductResponse, error){
        request := &GetProductRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetProduct(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Get the version information of the system.

         

         Returns
         -------
         False
         Version : Version
              Version information about the system

         
    */


    func(s *ServiceImpl)GetVersion() (*GetVersionResponse, error){
        request := &GetVersionRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetVersion(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Get the speed factor of a simulation (with lockstep a simulation can run faster or slower than realtime).

         

         Returns
         -------
         False
         SpeedFactor : float64
              Speed factor of simulation

         
    */


    func(s *ServiceImpl)GetSpeedFactor() (*GetSpeedFactorResponse, error){
        request := &GetSpeedFactorRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetSpeedFactor(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       