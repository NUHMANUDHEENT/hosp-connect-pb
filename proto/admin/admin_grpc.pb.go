// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package admin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	// SignIn method for admin users
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	// AddDoctor method for adding a new doctor
	AddDoctor(ctx context.Context, in *AddDoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error)
	// UpdateDoctor method for updating doctor details
	UpdateDoctor(ctx context.Context, in *UpdateDoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error)
	// DeleteDoctor method for deleting a doctor by ID
	DeleteDoctor(ctx context.Context, in *DeleteDoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error)
	// CreatePatient method for adding a new patient
	CreatePatient(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*PatientResponse, error)
	// UpdatePatient method for updating patient details
	UpdatePatient(ctx context.Context, in *UpdatePatientRequest, opts ...grpc.CallOption) (*PatientResponse, error)
	// DeletePatient method for deleting a patient by ID
	DeletePatient(ctx context.Context, in *DeletePatientRequest, opts ...grpc.CallOption) (*PatientResponse, error)
	// ViewStatistics method for retrieving hospital statistics
	ViewStatistics(ctx context.Context, in *ViewStatisticsRequest, opts ...grpc.CallOption) (*ViewStatisticsResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AddDoctor(ctx context.Context, in *AddDoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error) {
	out := new(DoctorResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/AddDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UpdateDoctor(ctx context.Context, in *UpdateDoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error) {
	out := new(DoctorResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/UpdateDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteDoctor(ctx context.Context, in *DeleteDoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error) {
	out := new(DoctorResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/DeleteDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreatePatient(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*PatientResponse, error) {
	out := new(PatientResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/CreatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UpdatePatient(ctx context.Context, in *UpdatePatientRequest, opts ...grpc.CallOption) (*PatientResponse, error) {
	out := new(PatientResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/UpdatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeletePatient(ctx context.Context, in *DeletePatientRequest, opts ...grpc.CallOption) (*PatientResponse, error) {
	out := new(PatientResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/DeletePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ViewStatistics(ctx context.Context, in *ViewStatisticsRequest, opts ...grpc.CallOption) (*ViewStatisticsResponse, error) {
	out := new(ViewStatisticsResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/ViewStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	// SignIn method for admin users
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	// AddDoctor method for adding a new doctor
	AddDoctor(context.Context, *AddDoctorRequest) (*DoctorResponse, error)
	// UpdateDoctor method for updating doctor details
	UpdateDoctor(context.Context, *UpdateDoctorRequest) (*DoctorResponse, error)
	// DeleteDoctor method for deleting a doctor by ID
	DeleteDoctor(context.Context, *DeleteDoctorRequest) (*DoctorResponse, error)
	// CreatePatient method for adding a new patient
	CreatePatient(context.Context, *CreatePatientRequest) (*PatientResponse, error)
	// UpdatePatient method for updating patient details
	UpdatePatient(context.Context, *UpdatePatientRequest) (*PatientResponse, error)
	// DeletePatient method for deleting a patient by ID
	DeletePatient(context.Context, *DeletePatientRequest) (*PatientResponse, error)
	// ViewStatistics method for retrieving hospital statistics
	ViewStatistics(context.Context, *ViewStatisticsRequest) (*ViewStatisticsResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAdminServiceServer) AddDoctor(context.Context, *AddDoctorRequest) (*DoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDoctor not implemented")
}
func (UnimplementedAdminServiceServer) UpdateDoctor(context.Context, *UpdateDoctorRequest) (*DoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoctor not implemented")
}
func (UnimplementedAdminServiceServer) DeleteDoctor(context.Context, *DeleteDoctorRequest) (*DoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDoctor not implemented")
}
func (UnimplementedAdminServiceServer) CreatePatient(context.Context, *CreatePatientRequest) (*PatientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatient not implemented")
}
func (UnimplementedAdminServiceServer) UpdatePatient(context.Context, *UpdatePatientRequest) (*PatientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatient not implemented")
}
func (UnimplementedAdminServiceServer) DeletePatient(context.Context, *DeletePatientRequest) (*PatientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePatient not implemented")
}
func (UnimplementedAdminServiceServer) ViewStatistics(context.Context, *ViewStatisticsRequest) (*ViewStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewStatistics not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AddDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/AddDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddDoctor(ctx, req.(*AddDoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UpdateDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UpdateDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/UpdateDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UpdateDoctor(ctx, req.(*UpdateDoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/DeleteDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteDoctor(ctx, req.(*DeleteDoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/CreatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreatePatient(ctx, req.(*CreatePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UpdatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UpdatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/UpdatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UpdatePatient(ctx, req.(*UpdatePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeletePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeletePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/DeletePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeletePatient(ctx, req.(*DeletePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ViewStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ViewStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/ViewStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ViewStatistics(ctx, req.(*ViewStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _AdminService_SignIn_Handler,
		},
		{
			MethodName: "AddDoctor",
			Handler:    _AdminService_AddDoctor_Handler,
		},
		{
			MethodName: "UpdateDoctor",
			Handler:    _AdminService_UpdateDoctor_Handler,
		},
		{
			MethodName: "DeleteDoctor",
			Handler:    _AdminService_DeleteDoctor_Handler,
		},
		{
			MethodName: "CreatePatient",
			Handler:    _AdminService_CreatePatient_Handler,
		},
		{
			MethodName: "UpdatePatient",
			Handler:    _AdminService_UpdatePatient_Handler,
		},
		{
			MethodName: "DeletePatient",
			Handler:    _AdminService_DeletePatient_Handler,
		},
		{
			MethodName: "ViewStatistics",
			Handler:    _AdminService_ViewStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/admin.proto",
}
