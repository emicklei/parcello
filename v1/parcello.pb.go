// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/parcello.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	v1/parcello.proto

It has these top-level messages:
	DeliverRequest
	DeliverResponse
	Envelope
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeliverRequest struct {
	Envelope *Envelope `protobuf:"bytes,1,opt,name=envelope" json:"envelope,omitempty"`
}

func (m *DeliverRequest) Reset()                    { *m = DeliverRequest{} }
func (m *DeliverRequest) String() string            { return proto.CompactTextString(m) }
func (*DeliverRequest) ProtoMessage()               {}
func (*DeliverRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeliverRequest) GetEnvelope() *Envelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

type DeliverResponse struct {
	ErrorMessage string `protobuf:"bytes,1,opt,name=errorMessage" json:"errorMessage,omitempty"`
}

func (m *DeliverResponse) Reset()                    { *m = DeliverResponse{} }
func (m *DeliverResponse) String() string            { return proto.CompactTextString(m) }
func (*DeliverResponse) ProtoMessage()               {}
func (*DeliverResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeliverResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type Envelope struct {
	// ID -> pubsub.Message.Attributes["parcello.ID"]
	// set by the server
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	// payload -> pubsub.Message.Data
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// attributes -> pubsub.Message.Attributes
	Attributes map[string]string `protobuf:"bytes,3,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// destinationTopic -> pubsub.Message.Attributes["parcello.destinationTopic"]
	// required for the DeliverRequest
	DestinationTopic string `protobuf:"bytes,4,opt,name=destinationTopic" json:"destinationTopic,omitempty"`
	// publishAfter -> pubsub.Message.Attributes["parcello.publishAfter"]
	// required for the DeliverRequest
	PublishAfter uint64 `protobuf:"varint,5,opt,name=publishAfter" json:"publishAfter,omitempty"`
	// deliveredAt -> pubsub.Message.Attributes["parcello.deliveredAt"]
	// set by the server to the time the DeliverRequest was sent
	DeliveredAt uint64 `protobuf:"varint,6,opt,name=deliveredAt" json:"deliveredAt,omitempty"`
	// publishCount -> pubsub.Message.Attributes["parcello.publishCount"]
	// set by the server to 1 on first publish to destination
	PublishCount uint64 `protobuf:"varint,7,opt,name=publishCount" json:"publishCount,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Envelope) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Envelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Envelope) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Envelope) GetDestinationTopic() string {
	if m != nil {
		return m.DestinationTopic
	}
	return ""
}

func (m *Envelope) GetPublishAfter() uint64 {
	if m != nil {
		return m.PublishAfter
	}
	return 0
}

func (m *Envelope) GetDeliveredAt() uint64 {
	if m != nil {
		return m.DeliveredAt
	}
	return 0
}

func (m *Envelope) GetPublishCount() uint64 {
	if m != nil {
		return m.PublishCount
	}
	return 0
}

func init() {
	proto.RegisterType((*DeliverRequest)(nil), "v1.DeliverRequest")
	proto.RegisterType((*DeliverResponse)(nil), "v1.DeliverResponse")
	proto.RegisterType((*Envelope)(nil), "v1.Envelope")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DeliveryService service

type DeliveryServiceClient interface {
	Deliver(ctx context.Context, in *DeliverRequest, opts ...grpc.CallOption) (*DeliverResponse, error)
}

type deliveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeliveryServiceClient(cc *grpc.ClientConn) DeliveryServiceClient {
	return &deliveryServiceClient{cc}
}

func (c *deliveryServiceClient) Deliver(ctx context.Context, in *DeliverRequest, opts ...grpc.CallOption) (*DeliverResponse, error) {
	out := new(DeliverResponse)
	err := grpc.Invoke(ctx, "/v1.DeliveryService/Deliver", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeliveryService service

type DeliveryServiceServer interface {
	Deliver(context.Context, *DeliverRequest) (*DeliverResponse, error)
}

func RegisterDeliveryServiceServer(s *grpc.Server, srv DeliveryServiceServer) {
	s.RegisterService(&_DeliveryService_serviceDesc, srv)
}

func _DeliveryService_Deliver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).Deliver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DeliveryService/Deliver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).Deliver(ctx, req.(*DeliverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeliveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DeliveryService",
	HandlerType: (*DeliveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deliver",
			Handler:    _DeliveryService_Deliver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/parcello.proto",
}

func init() { proto.RegisterFile("v1/parcello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4d, 0xcf, 0x93, 0x40,
	0x10, 0xc7, 0x03, 0x7d, 0x9e, 0xbe, 0x4c, 0x9b, 0xb6, 0xae, 0x1e, 0x36, 0x8d, 0x89, 0xc8, 0x89,
	0x78, 0xa0, 0x01, 0x63, 0x62, 0x1a, 0x3d, 0x54, 0xdb, 0x43, 0x0f, 0x5e, 0xd0, 0x93, 0x37, 0x5e,
	0xc6, 0x76, 0xd3, 0x2d, 0x8b, 0xbb, 0xcb, 0x26, 0x7c, 0x42, 0xbf, 0x96, 0x29, 0x50, 0x02, 0x7a,
	0x63, 0x7e, 0x33, 0xfc, 0x77, 0x66, 0xfe, 0x03, 0x2f, 0x4c, 0xb0, 0x2d, 0x62, 0x99, 0x22, 0xe7,
	0xc2, 0x2f, 0xa4, 0xd0, 0x82, 0xd8, 0x26, 0x70, 0x77, 0xb0, 0x3c, 0x20, 0x67, 0x06, 0x65, 0x84,
	0xbf, 0x4b, 0x54, 0x9a, 0x78, 0x30, 0xc5, 0xdc, 0x20, 0x17, 0x05, 0x52, 0xcb, 0xb1, 0xbc, 0x79,
	0xb8, 0xf0, 0x4d, 0xe0, 0x1f, 0x5b, 0x16, 0x75, 0x59, 0xf7, 0x03, 0xac, 0xba, 0x7f, 0x55, 0x21,
	0x72, 0x85, 0xc4, 0x85, 0x05, 0x4a, 0x29, 0xe4, 0x37, 0x54, 0x2a, 0x3e, 0x37, 0x02, 0xb3, 0x68,
	0xc0, 0xdc, 0x3f, 0x36, 0x4c, 0x1f, 0x6a, 0x64, 0x09, 0xf6, 0xe9, 0xd0, 0x96, 0xd9, 0xa7, 0x03,
	0xa1, 0x30, 0x29, 0xe2, 0x8a, 0x8b, 0x38, 0xa3, 0xb6, 0x63, 0x79, 0x8b, 0xe8, 0x11, 0x92, 0x4f,
	0x00, 0xb1, 0xd6, 0x92, 0x25, 0xa5, 0x46, 0x45, 0x47, 0xce, 0xc8, 0x9b, 0x87, 0xaf, 0xfb, 0x9d,
	0xf9, 0xfb, 0x2e, 0x7d, 0xcc, 0xb5, 0xac, 0xa2, 0x5e, 0x3d, 0x79, 0x07, 0xeb, 0x0c, 0x95, 0x66,
	0x79, 0xac, 0x99, 0xc8, 0x7f, 0x88, 0x82, 0xa5, 0xf4, 0xa9, 0x7e, 0xf5, 0x3f, 0x7e, 0x1f, 0xa2,
	0x28, 0x13, 0xce, 0xd4, 0x65, 0xff, 0x4b, 0xa3, 0xa4, 0xcf, 0x8e, 0xe5, 0x3d, 0x45, 0x03, 0x46,
	0x1c, 0x98, 0x67, 0xcd, 0xec, 0x98, 0xed, 0x35, 0x1d, 0xd7, 0x25, 0x7d, 0xd4, 0x53, 0xf9, 0x2a,
	0xca, 0x5c, 0xd3, 0xc9, 0x40, 0xa5, 0x66, 0x9b, 0xcf, 0xb0, 0xfa, 0xa7, 0x69, 0xb2, 0x86, 0xd1,
	0x15, 0xab, 0x76, 0x23, 0xf7, 0x4f, 0xf2, 0x0a, 0x9e, 0x4d, 0xcc, 0x4b, 0xac, 0x17, 0x32, 0x8b,
	0x9a, 0x60, 0x67, 0x7f, 0xb4, 0xc2, 0x63, 0x67, 0x40, 0xf5, 0x1d, 0xa5, 0x61, 0x29, 0x92, 0x10,
	0x26, 0x2d, 0x22, 0xe4, 0xbe, 0x9c, 0xa1, 0xb9, 0x9b, 0x97, 0x03, 0xd6, 0x98, 0xf6, 0xe5, 0xed,
	0xcf, 0x37, 0x67, 0xa6, 0x2f, 0x65, 0xe2, 0xa7, 0xe2, 0xb6, 0xc5, 0x1b, 0x4b, 0xaf, 0x1c, 0x59,
	0x77, 0x2d, 0x5b, 0x13, 0x24, 0xe3, 0xfa, 0x62, 0xde, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x0d,
	0x2a, 0x20, 0x49, 0x46, 0x02, 0x00, 0x00,
}
