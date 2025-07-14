package telemetry

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client TelemetryServiceClient
}

/*
Position Subscribe to 'position' updates.
*/
func (a *ServiceImpl) Position(
	ctx context.Context,

) (<-chan *Position, error) {
	ch := make(chan *Position)
	request := &SubscribePositionRequest{}
	stream, err := a.Client.SubscribePosition(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &PositionResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Position messages, err: %v", err)
			}
			ch <- m.GetPosition()
		}
	}()
	return ch, nil
}

/*
Home Subscribe to 'home position' updates.
*/
func (a *ServiceImpl) Home(
	ctx context.Context,

) (<-chan *Position, error) {
	ch := make(chan *Position)
	request := &SubscribeHomeRequest{}
	stream, err := a.Client.SubscribeHome(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &HomeResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Home messages, err: %v", err)
			}
			ch <- m.GetHome()
		}
	}()
	return ch, nil
}

/*
InAir Subscribe to in-air updates.
*/
func (a *ServiceImpl) InAir(
	ctx context.Context,

) (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeInAirRequest{}
	stream, err := a.Client.SubscribeInAir(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &InAirResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive InAir messages, err: %v", err)
			}
			ch <- m.GetIsInAir()
		}
	}()
	return ch, nil
}

/*
LandedState Subscribe to landed state updates
*/
func (a *ServiceImpl) LandedState(
	ctx context.Context,

) (<-chan LandedState, error) {
	ch := make(chan LandedState)
	request := &SubscribeLandedStateRequest{}
	stream, err := a.Client.SubscribeLandedState(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &LandedStateResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive LandedState messages, err: %v", err)
			}
			ch <- m.GetLandedState()
		}
	}()
	return ch, nil
}

/*
Armed Subscribe to armed updates.
*/
func (a *ServiceImpl) Armed(
	ctx context.Context,

) (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeArmedRequest{}
	stream, err := a.Client.SubscribeArmed(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ArmedResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Armed messages, err: %v", err)
			}
			ch <- m.GetIsArmed()
		}
	}()
	return ch, nil
}

/*
VtolState subscribe to vtol state Updates
*/
func (a *ServiceImpl) VtolState(
	ctx context.Context,

) (<-chan VtolState, error) {
	ch := make(chan VtolState)
	request := &SubscribeVtolStateRequest{}
	stream, err := a.Client.SubscribeVtolState(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &VtolStateResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive VtolState messages, err: %v", err)
			}
			ch <- m.GetVtolState()
		}
	}()
	return ch, nil
}

/*
AttitudeQuaternion Subscribe to 'attitude' updates (quaternion).
*/
func (a *ServiceImpl) AttitudeQuaternion(
	ctx context.Context,

) (<-chan *Quaternion, error) {
	ch := make(chan *Quaternion)
	request := &SubscribeAttitudeQuaternionRequest{}
	stream, err := a.Client.SubscribeAttitudeQuaternion(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &AttitudeQuaternionResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive AttitudeQuaternion messages, err: %v", err)
			}
			ch <- m.GetAttitudeQuaternion()
		}
	}()
	return ch, nil
}

/*
AttitudeEuler Subscribe to 'attitude' updates (Euler).
*/
func (a *ServiceImpl) AttitudeEuler(
	ctx context.Context,

) (<-chan *EulerAngle, error) {
	ch := make(chan *EulerAngle)
	request := &SubscribeAttitudeEulerRequest{}
	stream, err := a.Client.SubscribeAttitudeEuler(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &AttitudeEulerResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive AttitudeEuler messages, err: %v", err)
			}
			ch <- m.GetAttitudeEuler()
		}
	}()
	return ch, nil
}

/*
AttitudeAngularVelocityBody Subscribe to 'attitude' updates (angular velocity)
*/
func (a *ServiceImpl) AttitudeAngularVelocityBody(
	ctx context.Context,

) (<-chan *AngularVelocityBody, error) {
	ch := make(chan *AngularVelocityBody)
	request := &SubscribeAttitudeAngularVelocityBodyRequest{}
	stream, err := a.Client.SubscribeAttitudeAngularVelocityBody(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &AttitudeAngularVelocityBodyResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive AttitudeAngularVelocityBody messages, err: %v", err)
			}
			ch <- m.GetAttitudeAngularVelocityBody()
		}
	}()
	return ch, nil
}

/*
VelocityNed Subscribe to 'ground speed' updates (NED).
*/
func (a *ServiceImpl) VelocityNed(
	ctx context.Context,

) (<-chan *VelocityNed, error) {
	ch := make(chan *VelocityNed)
	request := &SubscribeVelocityNedRequest{}
	stream, err := a.Client.SubscribeVelocityNed(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &VelocityNedResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive VelocityNed messages, err: %v", err)
			}
			ch <- m.GetVelocityNed()
		}
	}()
	return ch, nil
}

/*
GpsInfo Subscribe to 'GPS info' updates.
*/
func (a *ServiceImpl) GpsInfo(
	ctx context.Context,

) (<-chan *GpsInfo, error) {
	ch := make(chan *GpsInfo)
	request := &SubscribeGpsInfoRequest{}
	stream, err := a.Client.SubscribeGpsInfo(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &GpsInfoResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive GpsInfo messages, err: %v", err)
			}
			ch <- m.GetGpsInfo()
		}
	}()
	return ch, nil
}

/*
RawGps Subscribe to 'Raw GPS' updates.
*/
func (a *ServiceImpl) RawGps(
	ctx context.Context,

) (<-chan *RawGps, error) {
	ch := make(chan *RawGps)
	request := &SubscribeRawGpsRequest{}
	stream, err := a.Client.SubscribeRawGps(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &RawGpsResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive RawGps messages, err: %v", err)
			}
			ch <- m.GetRawGps()
		}
	}()
	return ch, nil
}

/*
Battery Subscribe to 'battery' updates.
*/
func (a *ServiceImpl) Battery(
	ctx context.Context,

) (<-chan *Battery, error) {
	ch := make(chan *Battery)
	request := &SubscribeBatteryRequest{}
	stream, err := a.Client.SubscribeBattery(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &BatteryResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Battery messages, err: %v", err)
			}
			ch <- m.GetBattery()
		}
	}()
	return ch, nil
}

/*
FlightMode Subscribe to 'flight mode' updates.
*/
func (a *ServiceImpl) FlightMode(
	ctx context.Context,

) (<-chan FlightMode, error) {
	ch := make(chan FlightMode)
	request := &SubscribeFlightModeRequest{}
	stream, err := a.Client.SubscribeFlightMode(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &FlightModeResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive FlightMode messages, err: %v", err)
			}
			ch <- m.GetFlightMode()
		}
	}()
	return ch, nil
}

/*
Health Subscribe to 'health' updates.
*/
func (a *ServiceImpl) Health(
	ctx context.Context,

) (<-chan *Health, error) {
	ch := make(chan *Health)
	request := &SubscribeHealthRequest{}
	stream, err := a.Client.SubscribeHealth(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &HealthResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Health messages, err: %v", err)
			}
			ch <- m.GetHealth()
		}
	}()
	return ch, nil
}

/*
RcStatus Subscribe to 'RC status' updates.
*/
func (a *ServiceImpl) RcStatus(
	ctx context.Context,

) (<-chan *RcStatus, error) {
	ch := make(chan *RcStatus)
	request := &SubscribeRcStatusRequest{}
	stream, err := a.Client.SubscribeRcStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &RcStatusResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive RcStatus messages, err: %v", err)
			}
			ch <- m.GetRcStatus()
		}
	}()
	return ch, nil
}

/*
StatusText Subscribe to 'status text' updates.
*/
func (a *ServiceImpl) StatusText(
	ctx context.Context,

) (<-chan *StatusText, error) {
	ch := make(chan *StatusText)
	request := &SubscribeStatusTextRequest{}
	stream, err := a.Client.SubscribeStatusText(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &StatusTextResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive StatusText messages, err: %v", err)
			}
			ch <- m.GetStatusText()
		}
	}()
	return ch, nil
}

/*
ActuatorControlTarget Subscribe to 'actuator control target' updates.
*/
func (a *ServiceImpl) ActuatorControlTarget(
	ctx context.Context,

) (<-chan *ActuatorControlTarget, error) {
	ch := make(chan *ActuatorControlTarget)
	request := &SubscribeActuatorControlTargetRequest{}
	stream, err := a.Client.SubscribeActuatorControlTarget(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ActuatorControlTargetResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive ActuatorControlTarget messages, err: %v", err)
			}
			ch <- m.GetActuatorControlTarget()
		}
	}()
	return ch, nil
}

/*
ActuatorOutputStatus Subscribe to 'actuator output status' updates.
*/
func (a *ServiceImpl) ActuatorOutputStatus(
	ctx context.Context,

) (<-chan *ActuatorOutputStatus, error) {
	ch := make(chan *ActuatorOutputStatus)
	request := &SubscribeActuatorOutputStatusRequest{}
	stream, err := a.Client.SubscribeActuatorOutputStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ActuatorOutputStatusResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive ActuatorOutputStatus messages, err: %v", err)
			}
			ch <- m.GetActuatorOutputStatus()
		}
	}()
	return ch, nil
}

/*
Odometry Subscribe to 'odometry' updates.
*/
func (a *ServiceImpl) Odometry(
	ctx context.Context,

) (<-chan *Odometry, error) {
	ch := make(chan *Odometry)
	request := &SubscribeOdometryRequest{}
	stream, err := a.Client.SubscribeOdometry(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &OdometryResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Odometry messages, err: %v", err)
			}
			ch <- m.GetOdometry()
		}
	}()
	return ch, nil
}

/*
PositionVelocityNed Subscribe to 'position velocity' updates.
*/
func (a *ServiceImpl) PositionVelocityNed(
	ctx context.Context,

) (<-chan *PositionVelocityNed, error) {
	ch := make(chan *PositionVelocityNed)
	request := &SubscribePositionVelocityNedRequest{}
	stream, err := a.Client.SubscribePositionVelocityNed(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &PositionVelocityNedResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive PositionVelocityNed messages, err: %v", err)
			}
			ch <- m.GetPositionVelocityNed()
		}
	}()
	return ch, nil
}

/*
GroundTruth Subscribe to 'ground truth' updates.
*/
func (a *ServiceImpl) GroundTruth(
	ctx context.Context,

) (<-chan *GroundTruth, error) {
	ch := make(chan *GroundTruth)
	request := &SubscribeGroundTruthRequest{}
	stream, err := a.Client.SubscribeGroundTruth(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &GroundTruthResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive GroundTruth messages, err: %v", err)
			}
			ch <- m.GetGroundTruth()
		}
	}()
	return ch, nil
}

/*
FixedwingMetrics Subscribe to 'fixedwing metrics' updates.
*/
func (a *ServiceImpl) FixedwingMetrics(
	ctx context.Context,

) (<-chan *FixedwingMetrics, error) {
	ch := make(chan *FixedwingMetrics)
	request := &SubscribeFixedwingMetricsRequest{}
	stream, err := a.Client.SubscribeFixedwingMetrics(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &FixedwingMetricsResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive FixedwingMetrics messages, err: %v", err)
			}
			ch <- m.GetFixedwingMetrics()
		}
	}()
	return ch, nil
}

/*
Imu Subscribe to 'IMU' updates (in SI units in NED body frame).
*/
func (a *ServiceImpl) Imu(
	ctx context.Context,

) (<-chan *Imu, error) {
	ch := make(chan *Imu)
	request := &SubscribeImuRequest{}
	stream, err := a.Client.SubscribeImu(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ImuResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Imu messages, err: %v", err)
			}
			ch <- m.GetImu()
		}
	}()
	return ch, nil
}

/*
ScaledImu Subscribe to 'Scaled IMU' updates.
*/
func (a *ServiceImpl) ScaledImu(
	ctx context.Context,

) (<-chan *Imu, error) {
	ch := make(chan *Imu)
	request := &SubscribeScaledImuRequest{}
	stream, err := a.Client.SubscribeScaledImu(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ScaledImuResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive ScaledImu messages, err: %v", err)
			}
			ch <- m.GetImu()
		}
	}()
	return ch, nil
}

/*
RawImu Subscribe to 'Raw IMU' updates.
*/
func (a *ServiceImpl) RawImu(
	ctx context.Context,

) (<-chan *Imu, error) {
	ch := make(chan *Imu)
	request := &SubscribeRawImuRequest{}
	stream, err := a.Client.SubscribeRawImu(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &RawImuResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive RawImu messages, err: %v", err)
			}
			ch <- m.GetImu()
		}
	}()
	return ch, nil
}

/*
HealthAllOk Subscribe to 'HealthAllOk' updates.
*/
func (a *ServiceImpl) HealthAllOk(
	ctx context.Context,

) (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeHealthAllOkRequest{}
	stream, err := a.Client.SubscribeHealthAllOk(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &HealthAllOkResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive HealthAllOk messages, err: %v", err)
			}
			ch <- m.GetIsHealthAllOk()
		}
	}()
	return ch, nil
}

/*
UnixEpochTime Subscribe to 'unix epoch time' updates.
*/
func (a *ServiceImpl) UnixEpochTime(
	ctx context.Context,

) (<-chan uint64, error) {
	ch := make(chan uint64)
	request := &SubscribeUnixEpochTimeRequest{}
	stream, err := a.Client.SubscribeUnixEpochTime(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &UnixEpochTimeResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive UnixEpochTime messages, err: %v", err)
			}
			ch <- m.GetTimeUs()
		}
	}()
	return ch, nil
}

/*
DistanceSensor Subscribe to 'Distance Sensor' updates.
*/
func (a *ServiceImpl) DistanceSensor(
	ctx context.Context,

) (<-chan *DistanceSensor, error) {
	ch := make(chan *DistanceSensor)
	request := &SubscribeDistanceSensorRequest{}
	stream, err := a.Client.SubscribeDistanceSensor(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &DistanceSensorResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive DistanceSensor messages, err: %v", err)
			}
			ch <- m.GetDistanceSensor()
		}
	}()
	return ch, nil
}

/*
ScaledPressure Subscribe to 'Scaled Pressure' updates.
*/
func (a *ServiceImpl) ScaledPressure(
	ctx context.Context,

) (<-chan *ScaledPressure, error) {
	ch := make(chan *ScaledPressure)
	request := &SubscribeScaledPressureRequest{}
	stream, err := a.Client.SubscribeScaledPressure(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ScaledPressureResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive ScaledPressure messages, err: %v", err)
			}
			ch <- m.GetScaledPressure()
		}
	}()
	return ch, nil
}

/*
Heading Subscribe to 'Heading' updates.
*/
func (a *ServiceImpl) Heading(
	ctx context.Context,

) (<-chan *Heading, error) {
	ch := make(chan *Heading)
	request := &SubscribeHeadingRequest{}
	stream, err := a.Client.SubscribeHeading(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &HeadingResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Heading messages, err: %v", err)
			}
			ch <- m.GetHeadingDeg()
		}
	}()
	return ch, nil
}

/*
Altitude Subscribe to 'Altitude' updates.
*/
func (a *ServiceImpl) Altitude(
	ctx context.Context,

) (<-chan *Altitude, error) {
	ch := make(chan *Altitude)
	request := &SubscribeAltitudeRequest{}
	stream, err := a.Client.SubscribeAltitude(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &AltitudeResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Altitude messages, err: %v", err)
			}
			ch <- m.GetAltitude()
		}
	}()
	return ch, nil
}

/*
Wind Subscribe to 'Wind Estimated' updates.
*/
func (a *ServiceImpl) Wind(
	ctx context.Context,

) (<-chan *Wind, error) {
	ch := make(chan *Wind)
	request := &SubscribeWindRequest{}
	stream, err := a.Client.SubscribeWind(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &WindResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Wind messages, err: %v", err)
			}
			ch <- m.GetWind()
		}
	}()
	return ch, nil
}

/*
SetRatePosition Set rate to 'position' updates.
*/
func (s *ServiceImpl) SetRatePosition(
	ctx context.Context,
	rateHz float64,

) (*SetRatePositionResponse, error) {
	request := &SetRatePositionRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRatePosition(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateHome Set rate to 'home position' updates.
*/
func (s *ServiceImpl) SetRateHome(
	ctx context.Context,
	rateHz float64,

) (*SetRateHomeResponse, error) {
	request := &SetRateHomeRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateHome(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateInAir Set rate to in-air updates.
*/
func (s *ServiceImpl) SetRateInAir(
	ctx context.Context,
	rateHz float64,

) (*SetRateInAirResponse, error) {
	request := &SetRateInAirRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateInAir(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateLandedState Set rate to landed state updates
*/
func (s *ServiceImpl) SetRateLandedState(
	ctx context.Context,
	rateHz float64,

) (*SetRateLandedStateResponse, error) {
	request := &SetRateLandedStateRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateLandedState(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateVtolState Set rate to VTOL state updates
*/
func (s *ServiceImpl) SetRateVtolState(
	ctx context.Context,
	rateHz float64,

) (*SetRateVtolStateResponse, error) {
	request := &SetRateVtolStateRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateVtolState(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateAttitudeQuaternion Set rate to 'attitude euler angle' updates.
*/
func (s *ServiceImpl) SetRateAttitudeQuaternion(
	ctx context.Context,
	rateHz float64,

) (*SetRateAttitudeQuaternionResponse, error) {
	request := &SetRateAttitudeQuaternionRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateAttitudeQuaternion(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateAttitudeEuler Set rate to 'attitude quaternion' updates.
*/
func (s *ServiceImpl) SetRateAttitudeEuler(
	ctx context.Context,
	rateHz float64,

) (*SetRateAttitudeEulerResponse, error) {
	request := &SetRateAttitudeEulerRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateAttitudeEuler(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateVelocityNed Set rate of camera attitude updates.

	Set rate to 'ground speed' updates (NED).
*/
func (s *ServiceImpl) SetRateVelocityNed(
	ctx context.Context,
	rateHz float64,

) (*SetRateVelocityNedResponse, error) {
	request := &SetRateVelocityNedRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateVelocityNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateGpsInfo Set rate to 'GPS info' updates.
*/
func (s *ServiceImpl) SetRateGpsInfo(
	ctx context.Context,
	rateHz float64,

) (*SetRateGpsInfoResponse, error) {
	request := &SetRateGpsInfoRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateGpsInfo(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateBattery Set rate to 'battery' updates.
*/
func (s *ServiceImpl) SetRateBattery(
	ctx context.Context,
	rateHz float64,

) (*SetRateBatteryResponse, error) {
	request := &SetRateBatteryRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateBattery(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateRcStatus Set rate to 'RC status' updates.
*/
func (s *ServiceImpl) SetRateRcStatus(
	ctx context.Context,
	rateHz float64,

) (*SetRateRcStatusResponse, error) {
	request := &SetRateRcStatusRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateRcStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateActuatorControlTarget Set rate to 'actuator control target' updates.
*/
func (s *ServiceImpl) SetRateActuatorControlTarget(
	ctx context.Context,
	rateHz float64,

) (*SetRateActuatorControlTargetResponse, error) {
	request := &SetRateActuatorControlTargetRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateActuatorControlTarget(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateActuatorOutputStatus Set rate to 'actuator output status' updates.
*/
func (s *ServiceImpl) SetRateActuatorOutputStatus(
	ctx context.Context,
	rateHz float64,

) (*SetRateActuatorOutputStatusResponse, error) {
	request := &SetRateActuatorOutputStatusRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateActuatorOutputStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateOdometry Set rate to 'odometry' updates.
*/
func (s *ServiceImpl) SetRateOdometry(
	ctx context.Context,
	rateHz float64,

) (*SetRateOdometryResponse, error) {
	request := &SetRateOdometryRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateOdometry(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRatePositionVelocityNed Set rate to 'position velocity' updates.
*/
func (s *ServiceImpl) SetRatePositionVelocityNed(
	ctx context.Context,
	rateHz float64,

) (*SetRatePositionVelocityNedResponse, error) {
	request := &SetRatePositionVelocityNedRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRatePositionVelocityNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateGroundTruth Set rate to 'ground truth' updates.
*/
func (s *ServiceImpl) SetRateGroundTruth(
	ctx context.Context,
	rateHz float64,

) (*SetRateGroundTruthResponse, error) {
	request := &SetRateGroundTruthRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateGroundTruth(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateFixedwingMetrics Set rate to 'fixedwing metrics' updates.
*/
func (s *ServiceImpl) SetRateFixedwingMetrics(
	ctx context.Context,
	rateHz float64,

) (*SetRateFixedwingMetricsResponse, error) {
	request := &SetRateFixedwingMetricsRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateFixedwingMetrics(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateImu Set rate to 'IMU' updates.
*/
func (s *ServiceImpl) SetRateImu(
	ctx context.Context,
	rateHz float64,

) (*SetRateImuResponse, error) {
	request := &SetRateImuRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateImu(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateScaledImu Set rate to 'Scaled IMU' updates.
*/
func (s *ServiceImpl) SetRateScaledImu(
	ctx context.Context,
	rateHz float64,

) (*SetRateScaledImuResponse, error) {
	request := &SetRateScaledImuRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateScaledImu(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateRawImu Set rate to 'Raw IMU' updates.
*/
func (s *ServiceImpl) SetRateRawImu(
	ctx context.Context,
	rateHz float64,

) (*SetRateRawImuResponse, error) {
	request := &SetRateRawImuRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateRawImu(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateUnixEpochTime Set rate to 'unix epoch time' updates.
*/
func (s *ServiceImpl) SetRateUnixEpochTime(
	ctx context.Context,
	rateHz float64,

) (*SetRateUnixEpochTimeResponse, error) {
	request := &SetRateUnixEpochTimeRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateUnixEpochTime(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateDistanceSensor Set rate to 'Distance Sensor' updates.
*/
func (s *ServiceImpl) SetRateDistanceSensor(
	ctx context.Context,
	rateHz float64,

) (*SetRateDistanceSensorResponse, error) {
	request := &SetRateDistanceSensorRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateDistanceSensor(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateAltitude Set rate to 'Altitude' updates.
*/
func (s *ServiceImpl) SetRateAltitude(
	ctx context.Context,
	rateHz float64,

) (*SetRateAltitudeResponse, error) {
	request := &SetRateAltitudeRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateAltitude(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRateHealth Set rate to 'Health' updates.
*/
func (s *ServiceImpl) SetRateHealth(
	ctx context.Context,
	rateHz float64,

) (*SetRateHealthResponse, error) {
	request := &SetRateHealthRequest{
		RateHz: rateHz,
	}
	response, err := s.Client.SetRateHealth(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetGpsGlobalOrigin Get the GPS location of where the estimator has been initialized.
*/
func (s *ServiceImpl) GetGpsGlobalOrigin(
	ctx context.Context,

) (*GetGpsGlobalOriginResponse, error) {
	request := &GetGpsGlobalOriginRequest{}
	response, err := s.Client.GetGpsGlobalOrigin(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
