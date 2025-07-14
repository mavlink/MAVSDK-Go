package calibration

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client CalibrationServiceClient
}

     /*
         Perform gyro calibration.

         
    */

    func (a *ServiceImpl) CalibrateGyro(ctx context.Context, ) (<-chan  *ProgressData , error){
    		ch := make(chan  *ProgressData )
    		request := &SubscribeCalibrateGyroRequest{}
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
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive CalibrateGyro messages, err: %v\n", err)
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

    func (a *ServiceImpl) CalibrateAccelerometer(ctx context.Context, ) (<-chan  *ProgressData , error){
    		ch := make(chan  *ProgressData )
    		request := &SubscribeCalibrateAccelerometerRequest{}
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
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive CalibrateAccelerometer messages, err: %v\n", err)
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

    func (a *ServiceImpl) CalibrateMagnetometer(ctx context.Context, ) (<-chan  *ProgressData , error){
    		ch := make(chan  *ProgressData )
    		request := &SubscribeCalibrateMagnetometerRequest{}
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
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive CalibrateMagnetometer messages, err: %v\n", err)
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

    func (a *ServiceImpl) CalibrateLevelHorizon(ctx context.Context, ) (<-chan  *ProgressData , error){
    		ch := make(chan  *ProgressData )
    		request := &SubscribeCalibrateLevelHorizonRequest{}
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
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive CalibrateLevelHorizon messages, err: %v\n", err)
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

    func (a *ServiceImpl) CalibrateGimbalAccelerometer(ctx context.Context, ) (<-chan  *ProgressData , error){
    		ch := make(chan  *ProgressData )
    		request := &SubscribeCalibrateGimbalAccelerometerRequest{}
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
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive CalibrateGimbalAccelerometer messages, err: %v\n", err)
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

    func(s *ServiceImpl)Cancel(ctx context.Context, )(*CancelResponse, error){
        
        request := &CancelRequest{}
    	response, err := s.Client.Cancel(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       