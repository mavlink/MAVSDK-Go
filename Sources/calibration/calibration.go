package calibration

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client CalibrationServiceClient
}

/*
   Perform gyro calibration.


*/

func (a *ServiceImpl) CalibrateGyro() (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeCalibrateGyroRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeCalibrateGyro(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &CalibrateGyroResponse{}
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
   Perform accelerometer calibration.


*/

func (a *ServiceImpl) CalibrateAccelerometer() (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeCalibrateAccelerometerRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeCalibrateAccelerometer(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &CalibrateAccelerometerResponse{}
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
   Perform magnetometer calibration.


*/

func (a *ServiceImpl) CalibrateMagnetometer() (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeCalibrateMagnetometerRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeCalibrateMagnetometer(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &CalibrateMagnetometerResponse{}
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
   Perform board level horizon calibration.


*/

func (a *ServiceImpl) CalibrateLevelHorizon() (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeCalibrateLevelHorizonRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeCalibrateLevelHorizon(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &CalibrateLevelHorizonResponse{}
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
   Perform gimbal accelerometer calibration.


*/

func (a *ServiceImpl) CalibrateGimbalAccelerometer() (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeCalibrateGimbalAccelerometerRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeCalibrateGimbalAccelerometer(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &CalibrateGimbalAccelerometerResponse{}
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
   Cancel ongoing calibration process.


*/

func (s *ServiceImpl) Cancel() (*CancelResponse, error) {

	request := &CancelRequest{}
	ctx := context.Background()
	response, err := s.Client.Cancel(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
