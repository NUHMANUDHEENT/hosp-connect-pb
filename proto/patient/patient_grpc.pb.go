// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package patient

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PatientServiceClient is the client API for PatientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PatientServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	SignUpVerify(ctx context.Context, in *SignUpVerifyRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	AddPrescription(ctx context.Context, in *AddPrescriptionRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	GetPrescription(ctx context.Context, in *GetPrescriptionRequest, opts ...grpc.CallOption) (*GetPrescriptionResponse, error)
}

type patientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientServiceClient(cc grpc.ClientConnInterface) PatientServiceClient {
	return &patientServiceClient{cc}
}

func (c *patientServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/patient.PatientService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) SignUpVerify(ctx context.Context, in *SignUpVerifyRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/patient.PatientService/SignUpVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/patient.PatientService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, "/patient.PatientService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/patient.PatientService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) AddPrescription(ctx context.Context, in *AddPrescriptionRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/patient.PatientService/AddPrescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) GetPrescription(ctx context.Context, in *GetPrescriptionRequest, opts ...grpc.CallOption) (*GetPrescriptionResponse, error) {
	out := new(GetPrescriptionResponse)
	err := c.cc.Invoke(ctx, "/patient.PatientService/GetPrescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientServiceServer is the server API for PatientService service.
// All implementations must embed UnimplementedPatientServiceServer
// for forward compatibility
type PatientServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*StandardResponse, error)
	SignUpVerify(context.Context, *SignUpVerifyRequest) (*StandardResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*StandardResponse, error)
	AddPrescription(context.Context, *AddPrescriptionRequest) (*StandardResponse, error)
	GetPrescription(context.Context, *GetPrescriptionRequest) (*GetPrescriptionResponse, error)
	mustEmbedUnimplementedPatientServiceServer()
}

// UnimplementedPatientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPatientServiceServer struct {
}

func (UnimplementedPatientServiceServer) SignUp(context.Context, *SignUpRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedPatientServiceServer) SignUpVerify(context.Context, *SignUpVerifyRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUpVerify not implemented")
}
func (UnimplementedPatientServiceServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedPatientServiceServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedPatientServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedPatientServiceServer) AddPrescription(context.Context, *AddPrescriptionRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPrescription not implemented")
}
func (UnimplementedPatientServiceServer) GetPrescription(context.Context, *GetPrescriptionRequest) (*GetPrescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrescription not implemented")
}
func (UnimplementedPatientServiceServer) mustEmbedUnimplementedPatientServiceServer() {}

// UnsafePatientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PatientServiceServer will
// result in compilation errors.
type UnsafePatientServiceServer interface {
	mustEmbedUnimplementedPatientServiceServer()
}

func RegisterPatientServiceServer(s *grpc.Server, srv PatientServiceServer) {
	s.RegisterService(&_PatientService_serviceDesc, srv)
}

func _PatientService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patient.PatientService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_SignUpVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).SignUpVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patient.PatientService/SignUpVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).SignUpVerify(ctx, req.(*SignUpVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patient.PatientService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patient.PatientService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patient.PatientService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_AddPrescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPrescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).AddPrescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patient.PatientService/AddPrescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).AddPrescription(ctx, req.(*AddPrescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_GetPrescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).GetPrescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patient.PatientService/GetPrescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).GetPrescription(ctx, req.(*GetPrescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PatientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "patient.PatientService",
	HandlerType: (*PatientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _PatientService_SignUp_Handler,
		},
		{
			MethodName: "SignUpVerify",
			Handler:    _PatientService_SignUpVerify_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _PatientService_SignIn_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _PatientService_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _PatientService_UpdateProfile_Handler,
		},
		{
			MethodName: "AddPrescription",
			Handler:    _PatientService_AddPrescription_Handler,
		},
		{
			MethodName: "GetPrescription",
			Handler:    _PatientService_GetPrescription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "patient/patient.proto",
}
