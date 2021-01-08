package info

import (
	"context"
)

type ServiceImpl struct {
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

func (s *ServiceImpl) GetFlightInformation(ctx context.Context) (*GetFlightInformationResponse, error) {
	request := &GetFlightInformationRequest{}
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

func (s *ServiceImpl) GetIdentification(ctx context.Context) (*GetIdentificationResponse, error) {
	request := &GetIdentificationRequest{}
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

func (s *ServiceImpl) GetProduct(ctx context.Context) (*GetProductResponse, error) {
	request := &GetProductRequest{}
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

func (s *ServiceImpl) GetVersion(ctx context.Context) (*GetVersionResponse, error) {
	request := &GetVersionRequest{}
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

func (s *ServiceImpl) GetSpeedFactor(ctx context.Context) (*GetSpeedFactorResponse, error) {
	request := &GetSpeedFactorRequest{}
	response, err := s.Client.GetSpeedFactor(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}
