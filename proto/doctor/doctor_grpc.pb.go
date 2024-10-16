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
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	AddPrescription(ctx context.Context, in *AddPrescriptionRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	AddDoctor(ctx context.Context, in *AddDoctorRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	StoreAccessToken(ctx context.Context, in *StoreAccessTokenRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	ConfirmSchedule(ctx context.Context, in *ConfirmScheduleRequest, opts ...grpc.CallOption) (*ConfirmScheduleResponse, error)
	GetAvailability(ctx context.Context, in *GetAvailabilityRequest, opts ...grpc.CallOption) (*GetAvailabilityResponse, error)
	CheckAvailabilityByDoctorId(ctx context.Context, in *CheckAvailabilityByDoctorIdRequest, opts ...grpc.CallOption) (*CheckAvailabilityByDoctorIdResponse, error)
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

func (c *doctorServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/UpdateSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) AddPrescription(ctx context.Context, in *AddPrescriptionRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/AddPrescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) AddDoctor(ctx context.Context, in *AddDoctorRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/AddDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) StoreAccessToken(ctx context.Context, in *StoreAccessTokenRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/StoreAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/GetAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) ConfirmSchedule(ctx context.Context, in *ConfirmScheduleRequest, opts ...grpc.CallOption) (*ConfirmScheduleResponse, error) {
	out := new(ConfirmScheduleResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/ConfirmSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetAvailability(ctx context.Context, in *GetAvailabilityRequest, opts ...grpc.CallOption) (*GetAvailabilityResponse, error) {
	out := new(GetAvailabilityResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/GetAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) CheckAvailabilityByDoctorId(ctx context.Context, in *CheckAvailabilityByDoctorIdRequest, opts ...grpc.CallOption) (*CheckAvailabilityByDoctorIdResponse, error) {
	out := new(CheckAvailabilityByDoctorIdResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/CheckAvailabilityByDoctorId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations must embed UnimplementedDoctorServiceServer
// for forward compatibility
type DoctorServiceServer interface {
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*StandardResponse, error)
	UpdateSchedule(context.Context, *UpdateScheduleRequest) (*StandardResponse, error)
	AddPrescription(context.Context, *AddPrescriptionRequest) (*StandardResponse, error)
	AddDoctor(context.Context, *AddDoctorRequest) (*StandardResponse, error)
	StoreAccessToken(context.Context, *StoreAccessTokenRequest) (*StandardResponse, error)
	GetAccessToken(context.Context, *GetAccessTokenRequest) (*StandardResponse, error)
	ConfirmSchedule(context.Context, *ConfirmScheduleRequest) (*ConfirmScheduleResponse, error)
	GetAvailability(context.Context, *GetAvailabilityRequest) (*GetAvailabilityResponse, error)
	CheckAvailabilityByDoctorId(context.Context, *CheckAvailabilityByDoctorIdRequest) (*CheckAvailabilityByDoctorIdResponse, error)
	mustEmbedUnimplementedDoctorServiceServer()
}

// UnimplementedDoctorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorServiceServer struct {
}

func (UnimplementedDoctorServiceServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedDoctorServiceServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedDoctorServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedDoctorServiceServer) UpdateSchedule(context.Context, *UpdateScheduleRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSchedule not implemented")
}
func (UnimplementedDoctorServiceServer) AddPrescription(context.Context, *AddPrescriptionRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPrescription not implemented")
}
func (UnimplementedDoctorServiceServer) AddDoctor(context.Context, *AddDoctorRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) StoreAccessToken(context.Context, *StoreAccessTokenRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreAccessToken not implemented")
}
func (UnimplementedDoctorServiceServer) GetAccessToken(context.Context, *GetAccessTokenRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedDoctorServiceServer) ConfirmSchedule(context.Context, *ConfirmScheduleRequest) (*ConfirmScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmSchedule not implemented")
}
func (UnimplementedDoctorServiceServer) GetAvailability(context.Context, *GetAvailabilityRequest) (*GetAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailability not implemented")
}
func (UnimplementedDoctorServiceServer) CheckAvailabilityByDoctorId(context.Context, *CheckAvailabilityByDoctorIdRequest) (*CheckAvailabilityByDoctorIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAvailabilityByDoctorId not implemented")
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

func _DoctorService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
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

func _DoctorService_StoreAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).StoreAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/StoreAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).StoreAccessToken(ctx, req.(*StoreAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/GetAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetAccessToken(ctx, req.(*GetAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_ConfirmSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).ConfirmSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/ConfirmSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).ConfirmSchedule(ctx, req.(*ConfirmScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/GetAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetAvailability(ctx, req.(*GetAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_CheckAvailabilityByDoctorId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAvailabilityByDoctorIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).CheckAvailabilityByDoctorId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/CheckAvailabilityByDoctorId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).CheckAvailabilityByDoctorId(ctx, req.(*CheckAvailabilityByDoctorIdRequest))
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
			MethodName: "GetProfile",
			Handler:    _DoctorService_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _DoctorService_UpdateProfile_Handler,
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
			MethodName: "AddDoctor",
			Handler:    _DoctorService_AddDoctor_Handler,
		},
		{
			MethodName: "StoreAccessToken",
			Handler:    _DoctorService_StoreAccessToken_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _DoctorService_GetAccessToken_Handler,
		},
		{
			MethodName: "ConfirmSchedule",
			Handler:    _DoctorService_ConfirmSchedule_Handler,
		},
		{
			MethodName: "GetAvailability",
			Handler:    _DoctorService_GetAvailability_Handler,
		},
		{
			MethodName: "CheckAvailabilityByDoctorId",
			Handler:    _DoctorService_CheckAvailabilityByDoctorId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doctor/doctor.proto",
}
