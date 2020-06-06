package gorpc_demo

import "time"

//ServerOptions 是通过选项模式来透传业务自己指定的一些参数
//比如服务监听的地址 address，
//网络类型 network 是 tcp 还是 udp，
//后端服务的超时时间 timeout 等。
type ServerOptions struct {
	address string  // listening address, e.g. :( ip://127.0.0.1:8080、 dns://www.google.com)
	network string  // network type, e.g. : tcp、udp
	protocol string  // protocol type, e.g. : proto、json
	timeout time.Duration       // timeout
	serializationType string 	// serialization type, default: proto

	selectorSvrAddr string       // service discovery server address, required when using the third-party service discovery plugin
	tracingSvrAddr  string 		 // tracing plugin server address, required when using the third-party tracing plugin
	tracingSpanName string       // tracing span name, required when using the third-party tracing plugin
	pluginNames []string         // plugin name
}

type ServerOption func(*ServerOptions)

func WithAddress(address string) ServerOption{
	return func(o *ServerOptions) {
		o.address = address
	}
}

func WithNetwork(network string) ServerOption {
	return func(o *ServerOptions) {
		o.network = network
	}
}

func WithProtocol(protocol string) ServerOption {
	return func(o *ServerOptions) {
		o.protocol = protocol
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(o *ServerOptions) {
		o.timeout = timeout
	}
}

func WithSerializationType(serializationType string) ServerOption {
	return func(o *ServerOptions) {
		o.serializationType = serializationType
	}
}

func WithSelectorSvrAddr(addr string) ServerOption {
	return func(o *ServerOptions) {
		o.selectorSvrAddr = addr
	}
}

func WithPlugin(pluginName ... string) ServerOption {
	return func(o *ServerOptions) {
		o.pluginNames = append(o.pluginNames, pluginName ...)
	}
}


func WithTracingSvrAddr(addr string) ServerOption {
	return func(o *ServerOptions) {
		o.tracingSvrAddr = addr
	}
}

func WithTracingSpanName(name string) ServerOption {
	return func(o *ServerOptions) {
		o.tracingSpanName = name
	}
}