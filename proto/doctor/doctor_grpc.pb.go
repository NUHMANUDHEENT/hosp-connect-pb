// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package doctor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DoctorServiceClient is the client API for DoctorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorServiceClient interface {
	// Doctor SignIn
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	// Update doctor profile details
	UpdateDoctorProfile(ctx context.Context, in *UpdateDoctorProfileRequest, opts ...grpc.CallOption) (*UpdateDoctorProfileResponse, error)
	// Update doctor schedule for one week with 30-minute slots
	UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*UpdateScheduleResponse, error)
	// Add a prescription for a patient
	AddPrescription(ctx context.Context, in *AddPrescriptionRequest, opts ...grpc.CallOption) (*AddPrescriptionResponse, error)
	// Get past prescriptions for a specific patient
	GetPastPrescriptions(ctx context.Context, in *GetPastPrescriptionsRequest, opts ...grpc.CallOption) (*GetPastPrescriptionsResponse, error)
	AddDoctor(ctx context.Context, in *AddDoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error)
}

type doctorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServiceClient(cc grpc.ClientConnInterface) DoctorServiceClient {
	return &doctorServiceClient{cc}
}

func (c *doctorServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) UpdateDoctorProfile(ctx context.Context, in *UpdateDoctorProfileRequest, opts ...grpc.CallOption) (*UpdateDoctorProfileResponse, error) {
	out := new(UpdateDoctorProfileResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/UpdateDoctorProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*UpdateScheduleResponse, error) {
	out := new(UpdateScheduleResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/UpdateSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) AddPrescription(ctx context.Context, in *AddPrescriptionRequest, opts ...grpc.CallOption) (*AddPrescriptionResponse, error) {
	out := new(AddPrescriptionResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/AddPrescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetPastPrescriptions(ctx context.Context, in *GetPastPrescriptionsRequest, opts ...grpc.CallOption) (*GetPastPrescriptionsResponse, error) {
	out := new(GetPastPrescriptionsResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/GetPastPrescriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) AddDoctor(ctx context.Context, in *AddDoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error) {
	out := new(DoctorResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/AddDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations must embed UnimplementedDoctorServiceServer
// for forward compatibility
type DoctorServiceServer interface {
	// Doctor SignIn
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	// Update doctor profile details
	UpdateDoctorProfile(context.Context, *UpdateDoctorProfileRequest) (*UpdateDoctorProfileResponse, error)
	// Update doctor schedule for one week with 30-minute slots
	UpdateSchedule(context.Context, *UpdateScheduleRequest) (*UpdateScheduleResponse, error)
	// Add a prescription for a patient
	AddPrescription(context.Context, *AddPrescriptionRequest) (*AddPrescriptionResponse, error)
	// Get past prescriptions for a specific patient
	GetPastPrescriptions(context.Context, *GetPastPrescriptionsRequest) (*GetPastPrescriptionsResponse, error)
	AddDoctor(context.Context, *AddDoctorRequest) (*DoctorResponse, error)
	mustEmbedUnimplementedDoctorServiceServer()
}

// UnimplementedDoctorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorServiceServer struct {
}

func (UnimplementedDoctorServiceServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedDoctorServiceServer) UpdateDoctorProfile(context.Context, *UpdateDoctorProfileRequest) (*UpdateDoctorProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoctorProfile not implemented")
}
func (UnimplementedDoctorServiceServer) UpdateSchedule(context.Context, *UpdateScheduleRequest) (*UpdateScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSchedule not implemented")
}
func (UnimplementedDoctorServiceServer) AddPrescription(context.Context, *AddPrescriptionRequest) (*AddPrescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPrescription not implemented")
}
func (UnimplementedDoctorServiceServer) GetPastPrescriptions(context.Context, *GetPastPrescriptionsRequest) (*GetPastPrescriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPastPrescriptions not implemented")
}
func (UnimplementedDoctorServiceServer) AddDoctor(context.Context, *AddDoctorRequest) (*DoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) mustEmbedUnimplementedDoctorServiceServer() {}

// UnsafeDoctorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServiceServer will
// result in compilation errors.
type UnsafeDoctorServiceServer interface {
	mustEmbedUnimplementedDoctorServiceServer()
}

func RegisterDoctorServiceServer(s *grpc.Server, srv DoctorServiceServer) {
	s.RegisterService(&_DoctorService_serviceDesc, srv)
}

func _DoctorService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_UpdateDoctorProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDoctorProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).UpdateDoctorProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/UpdateDoctorProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).UpdateDoctorProfile(ctx, req.(*UpdateDoctorProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_UpdateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).UpdateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/UpdateSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).UpdateSchedule(ctx, req.(*UpdateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_AddPrescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPrescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).AddPrescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/AddPrescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).AddPrescription(ctx, req.(*AddPrescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetPastPrescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPastPrescriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetPastPrescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/GetPastPrescriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetPastPrescriptions(ctx, req.(*GetPastPrescriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_AddDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).AddDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/AddDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).AddDoctor(ctx, req.(*AddDoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DoctorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "doctor.DoctorService",
	HandlerType: (*DoctorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _DoctorService_SignIn_Handler,
		},
		{
			MethodName: "UpdateDoctorProfile",
			Handler:    _DoctorService_UpdateDoctorProfile_Handler,
		},
		{
			MethodName: "UpdateSchedule",
			Handler:    _DoctorService_UpdateSchedule_Handler,
		},
		{
			MethodName: "AddPrescription",
			Handler:    _DoctorService_AddPrescription_Handler,
		},
		{
			MethodName: "GetPastPrescriptions",
			Handler:    _DoctorService_GetPastPrescriptions_Handler,
		},
		{
			MethodName: "AddDoctor",
			Handler:    _DoctorService_AddDoctor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doctor/doctor.proto",
}