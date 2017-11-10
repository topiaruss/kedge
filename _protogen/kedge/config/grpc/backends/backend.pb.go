// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kedge/config/grpc/backends/backend.proto

/*
Package kedge_config_grpc_backends is a generated protocol buffer package.

It is generated from these files:
	kedge/config/grpc/backends/backend.proto

It has these top-level messages:
	Backend
	Interceptor
	Security
*/
package kedge_config_grpc_backends

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"
import  kedge_config_common_resolvers "github.com/improbable-eng/kedge/_protogen/kedge/config/common/resolvers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// / Balancer chooses which gRPC balancing policy to use.
type Balancer int32

const (
	// ROUND_ROBIN is the simpliest and default load balancing policy
	Balancer_ROUND_ROBIN Balancer = 0
)

var Balancer_name = map[int32]string{
	0: "ROUND_ROBIN",
}
var Balancer_value = map[string]int32{
	"ROUND_ROBIN": 0,
}

func (x Balancer) String() string {
	return proto.EnumName(Balancer_name, int32(x))
}
func (Balancer) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// / Backend is a gRPC ClientConn pool maintained to a single serivce.
type Backend struct {
	// / name is the string identifying the bakcend in all other conifgs.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// / balancer decides which balancing policy to use.
	Balancer Balancer `protobuf:"varint,2,opt,name=balancer,enum=kedge.config.grpc.backends.Balancer" json:"balancer,omitempty"`
	// / disable_conntracking turns off the /debug/events tracing and Prometheus monitoring of the pool sie for this backend.
	DisableConntracking bool `protobuf:"varint,3,opt,name=disable_conntracking,json=disableConntracking" json:"disable_conntracking,omitempty"`
	// / security controls the TLS connection details for the backend. If not present, Insecure (plain text) mode is used.
	Security *Security `protobuf:"bytes,4,opt,name=security" json:"security,omitempty"`
	// / interceptors controls what interceptors will be enabled for this backend.
	Interceptors []*Interceptor `protobuf:"bytes,5,rep,name=interceptors" json:"interceptors,omitempty"`
	// Types that are valid to be assigned to Resolver:
	//	*Backend_Srv
	//	*Backend_K8S
	//	*Backend_Host
	Resolver      isBackend_Resolver `protobuf_oneof:"resolver"`
	Autogenerated bool               `protobuf:"varint,6,opt,name=autogenerated" json:"autogenerated,omitempty"`
}

func (m *Backend) Reset()                    { *m = Backend{} }
func (m *Backend) String() string            { return proto.CompactTextString(m) }
func (*Backend) ProtoMessage()               {}
func (*Backend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isBackend_Resolver interface {
	isBackend_Resolver()
}

type Backend_Srv struct {
	Srv *kedge_config_common_resolvers.SrvResolver `protobuf:"bytes,10,opt,name=srv,oneof"`
}
type Backend_K8S struct {
	K8S *kedge_config_common_resolvers.K8SResolver `protobuf:"bytes,11,opt,name=k8s,oneof"`
}
type Backend_Host struct {
	Host *kedge_config_common_resolvers.HostResolver `protobuf:"bytes,12,opt,name=host,oneof"`
}

func (*Backend_Srv) isBackend_Resolver()  {}
func (*Backend_K8S) isBackend_Resolver()  {}
func (*Backend_Host) isBackend_Resolver() {}

func (m *Backend) GetResolver() isBackend_Resolver {
	if m != nil {
		return m.Resolver
	}
	return nil
}

func (m *Backend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Backend) GetBalancer() Balancer {
	if m != nil {
		return m.Balancer
	}
	return Balancer_ROUND_ROBIN
}

func (m *Backend) GetDisableConntracking() bool {
	if m != nil {
		return m.DisableConntracking
	}
	return false
}

func (m *Backend) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *Backend) GetInterceptors() []*Interceptor {
	if m != nil {
		return m.Interceptors
	}
	return nil
}

func (m *Backend) GetSrv() *kedge_config_common_resolvers.SrvResolver {
	if x, ok := m.GetResolver().(*Backend_Srv); ok {
		return x.Srv
	}
	return nil
}

func (m *Backend) GetK8S() *kedge_config_common_resolvers.K8SResolver {
	if x, ok := m.GetResolver().(*Backend_K8S); ok {
		return x.K8S
	}
	return nil
}

func (m *Backend) GetHost() *kedge_config_common_resolvers.HostResolver {
	if x, ok := m.GetResolver().(*Backend_Host); ok {
		return x.Host
	}
	return nil
}

func (m *Backend) GetAutogenerated() bool {
	if m != nil {
		return m.Autogenerated
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Backend) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Backend_OneofMarshaler, _Backend_OneofUnmarshaler, _Backend_OneofSizer, []interface{}{
		(*Backend_Srv)(nil),
		(*Backend_K8S)(nil),
		(*Backend_Host)(nil),
	}
}

func _Backend_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Backend)
	// resolver
	switch x := m.Resolver.(type) {
	case *Backend_Srv:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Srv); err != nil {
			return err
		}
	case *Backend_K8S:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.K8S); err != nil {
			return err
		}
	case *Backend_Host:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Host); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Backend.Resolver has unexpected type %T", x)
	}
	return nil
}

func _Backend_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Backend)
	switch tag {
	case 10: // resolver.srv
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kedge_config_common_resolvers.SrvResolver)
		err := b.DecodeMessage(msg)
		m.Resolver = &Backend_Srv{msg}
		return true, err
	case 11: // resolver.k8s
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kedge_config_common_resolvers.K8SResolver)
		err := b.DecodeMessage(msg)
		m.Resolver = &Backend_K8S{msg}
		return true, err
	case 12: // resolver.host
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kedge_config_common_resolvers.HostResolver)
		err := b.DecodeMessage(msg)
		m.Resolver = &Backend_Host{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Backend_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Backend)
	// resolver
	switch x := m.Resolver.(type) {
	case *Backend_Srv:
		s := proto.Size(x.Srv)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Backend_K8S:
		s := proto.Size(x.K8S)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Backend_Host:
		s := proto.Size(x.Host)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Interceptor struct {
	// Types that are valid to be assigned to Interceptor:
	//	*Interceptor_Prometheus
	Interceptor isInterceptor_Interceptor `protobuf_oneof:"interceptor"`
}

func (m *Interceptor) Reset()                    { *m = Interceptor{} }
func (m *Interceptor) String() string            { return proto.CompactTextString(m) }
func (*Interceptor) ProtoMessage()               {}
func (*Interceptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isInterceptor_Interceptor interface {
	isInterceptor_Interceptor()
}

type Interceptor_Prometheus struct {
	Prometheus bool `protobuf:"varint,1,opt,name=prometheus,oneof"`
}

func (*Interceptor_Prometheus) isInterceptor_Interceptor() {}

func (m *Interceptor) GetInterceptor() isInterceptor_Interceptor {
	if m != nil {
		return m.Interceptor
	}
	return nil
}

func (m *Interceptor) GetPrometheus() bool {
	if x, ok := m.GetInterceptor().(*Interceptor_Prometheus); ok {
		return x.Prometheus
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Interceptor) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Interceptor_OneofMarshaler, _Interceptor_OneofUnmarshaler, _Interceptor_OneofSizer, []interface{}{
		(*Interceptor_Prometheus)(nil),
	}
}

func _Interceptor_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Interceptor)
	// interceptor
	switch x := m.Interceptor.(type) {
	case *Interceptor_Prometheus:
		t := uint64(0)
		if x.Prometheus {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("Interceptor.Interceptor has unexpected type %T", x)
	}
	return nil
}

func _Interceptor_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Interceptor)
	switch tag {
	case 1: // interceptor.prometheus
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Interceptor = &Interceptor_Prometheus{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _Interceptor_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Interceptor)
	// interceptor
	switch x := m.Interceptor.(type) {
	case *Interceptor_Prometheus:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// / Security settings for a backend.
type Security struct {
	// / insecure_skip_verify skips the server certificate verification completely.
	// / No TLS config (for testclient or server) will be used. This should *not* be used in production software.
	InsecureSkipVerify bool `protobuf:"varint,1,opt,name=insecure_skip_verify,json=insecureSkipVerify" json:"insecure_skip_verify,omitempty"`
	// / config_name indicates the TlsServerConfig to be used for this connection.
	ConfigName string `protobuf:"bytes,2,opt,name=config_name,json=configName" json:"config_name,omitempty"`
}

func (m *Security) Reset()                    { *m = Security{} }
func (m *Security) String() string            { return proto.CompactTextString(m) }
func (*Security) ProtoMessage()               {}
func (*Security) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Security) GetInsecureSkipVerify() bool {
	if m != nil {
		return m.InsecureSkipVerify
	}
	return false
}

func (m *Security) GetConfigName() string {
	if m != nil {
		return m.ConfigName
	}
	return ""
}

func init() {
	proto.RegisterType((*Backend)(nil), "kedge.config.grpc.backends.Backend")
	proto.RegisterType((*Interceptor)(nil), "kedge.config.grpc.backends.Interceptor")
	proto.RegisterType((*Security)(nil), "kedge.config.grpc.backends.Security")
	proto.RegisterEnum("kedge.config.grpc.backends.Balancer", Balancer_name, Balancer_value)
}

func init() { proto.RegisterFile("kedge/config/grpc/backends/backend.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x6d, 0x68, 0x28, 0xee, 0xb8, 0x05, 0xb4, 0xe4, 0x60, 0x85, 0x43, 0xad, 0x28, 0x12, 0x56,
	0x69, 0xec, 0x12, 0x50, 0x15, 0x2e, 0x15, 0x18, 0x0e, 0xa9, 0x2a, 0xa5, 0xd2, 0x46, 0x70, 0x41,
	0xc5, 0xda, 0x38, 0x5b, 0x67, 0xe5, 0x78, 0xd7, 0xda, 0xdd, 0xa4, 0x2a, 0x88, 0x7f, 0xe2, 0x8f,
	0x90, 0xf8, 0x12, 0xe4, 0x75, 0x9c, 0x38, 0x07, 0x0a, 0xb7, 0xf1, 0xcc, 0x7b, 0x6f, 0xbc, 0x6f,
	0x1e, 0x78, 0x29, 0x9d, 0x26, 0x34, 0x88, 0x05, 0xbf, 0x61, 0x49, 0x90, 0xc8, 0x3c, 0x0e, 0x26,
	0x24, 0x4e, 0x29, 0x9f, 0xaa, 0xaa, 0xf0, 0x73, 0x29, 0xb4, 0x40, 0x6d, 0x83, 0xf4, 0x4b, 0xa4,
	0x5f, 0x20, 0xfd, 0x0a, 0xd9, 0x3e, 0x4b, 0x98, 0x9e, 0x2d, 0x26, 0x7e, 0x2c, 0xb2, 0x20, 0xbb,
	0x65, 0x3a, 0x15, 0xb7, 0x41, 0x22, 0x7a, 0x86, 0xd8, 0x5b, 0x92, 0x39, 0x9b, 0x12, 0x2d, 0xa4,
	0x0a, 0xd6, 0x65, 0xa9, 0xd9, 0xee, 0x6d, 0x6d, 0x8f, 0x45, 0x96, 0x09, 0x1e, 0x48, 0xaa, 0xc4,
	0x7c, 0x49, 0xa5, 0xda, 0x54, 0x25, 0xbc, 0xf3, 0xb3, 0x09, 0x8f, 0xc2, 0x72, 0x27, 0x3a, 0x81,
	0x26, 0x27, 0x19, 0x75, 0x1a, 0x6e, 0xc3, 0xdb, 0x0f, 0x9d, 0xdf, 0xbf, 0x8e, 0x5a, 0x80, 0xbe,
	0x7e, 0x21, 0xbd, 0x6f, 0xd1, 0x69, 0xef, 0xad, 0x7f, 0xfd, 0xbd, 0x7f, 0x72, 0xf6, 0xe6, 0x47,
	0x17, 0x1b, 0x14, 0x7a, 0x07, 0xd6, 0x84, 0xcc, 0x09, 0x8f, 0xa9, 0x74, 0x1e, 0xb8, 0x0d, 0xef,
	0x71, 0xbf, 0xeb, 0xff, 0xfd, 0x3d, 0x7e, 0xb8, 0xc2, 0xe2, 0x35, 0x0b, 0xbd, 0x82, 0xd6, 0x94,
	0x29, 0x32, 0x99, 0xd3, 0x28, 0x16, 0x9c, 0x6b, 0x49, 0xe2, 0x94, 0xf1, 0xc4, 0xd9, 0x75, 0x1b,
	0x9e, 0x85, 0x9f, 0xad, 0x66, 0x1f, 0x6a, 0xa3, 0x62, 0xa9, 0xa2, 0xf1, 0x42, 0x32, 0x7d, 0xe7,
	0x34, 0xdd, 0x86, 0x67, 0xdf, 0xbf, 0x74, 0xbc, 0xc2, 0xe2, 0x35, 0x0b, 0x5d, 0xc2, 0x01, 0xe3,
	0x9a, 0xca, 0x98, 0xe6, 0x85, 0x7f, 0xce, 0x43, 0x77, 0xd7, 0xb3, 0xfb, 0x2f, 0xee, 0x53, 0xb9,
	0xd8, 0xe0, 0xf1, 0x16, 0x19, 0x9d, 0xc3, 0xae, 0x92, 0x4b, 0x07, 0xcc, 0x9f, 0x1c, 0x6f, 0x6b,
	0x94, 0xd6, 0xfb, 0x1b, 0xc3, 0xc7, 0x72, 0x89, 0x57, 0x1f, 0xc3, 0x1d, 0x5c, 0x10, 0x0b, 0x7e,
	0x3a, 0x50, 0x8e, 0xfd, 0x5f, 0xfc, 0xcb, 0x81, 0xaa, 0xf3, 0xd3, 0x81, 0x42, 0xef, 0xa1, 0x39,
	0x13, 0x4a, 0x3b, 0x07, 0x46, 0xe0, 0xe5, 0x3f, 0x04, 0x86, 0x42, 0xe9, 0x9a, 0x82, 0xa1, 0xa2,
	0x2e, 0x1c, 0x92, 0x85, 0x16, 0x09, 0xe5, 0x54, 0x12, 0x4d, 0xa7, 0xce, 0x9e, 0x71, 0x7f, 0xbb,
	0x19, 0x02, 0x58, 0x95, 0x4e, 0xe7, 0x1c, 0xec, 0x9a, 0x23, 0xc8, 0x05, 0xc8, 0xa5, 0xc8, 0xa8,
	0x9e, 0xd1, 0x85, 0x32, 0xd9, 0xb1, 0x86, 0x3b, 0xb8, 0xd6, 0x0b, 0x0f, 0xc1, 0xae, 0xb9, 0xd6,
	0xb9, 0x06, 0xab, 0xba, 0x0b, 0x3a, 0x85, 0x16, 0xe3, 0xe6, 0x36, 0x34, 0x52, 0x29, 0xcb, 0xa3,
	0x25, 0x95, 0xec, 0xe6, 0xae, 0x94, 0xc1, 0xa8, 0x9a, 0x8d, 0x53, 0x96, 0x7f, 0x36, 0x13, 0x74,
	0x04, 0x76, 0xf9, 0xbe, 0xc8, 0x64, 0xb5, 0x48, 0xde, 0x3e, 0x86, 0xb2, 0x35, 0x22, 0x19, 0x3d,
	0x7e, 0x0e, 0x56, 0x95, 0x35, 0xf4, 0x04, 0x6c, 0x7c, 0xf5, 0x69, 0xf4, 0x31, 0xc2, 0x57, 0xe1,
	0xc5, 0xe8, 0xe9, 0xce, 0x64, 0xcf, 0xa4, 0xfe, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19,
	0xe3, 0xcd, 0x6e, 0xa4, 0x03, 0x00, 0x00,
}
