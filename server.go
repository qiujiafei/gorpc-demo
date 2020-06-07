package gorpc_demo

import (
	"context"
	"fmt"
	"log"
	"reflect"
)

//server 定义服务，发布服务，
//接收到服务的请求后，根据服务名和请求的方法名去路由到一个 handler 处理器，
//然后由 handler 处理请求，得到响应，并且把响应数据发送给 client
type Server struct {
	opts *ServerOptions
	service Service
	closing bool
}

func NewServer(opt ...ServerOption) *Server {
	// 先初始化一个空的 opts
	s := &Server{
		opts: &ServerOptions{},
	}

	// 再往 空的 opts 里面赋值
	for _, o := range opt {
		o(s.opts)
	}

	s.service = NewService(s.opts)
	return s
}

func NewService(opts *ServerOptions) Service {
	return &service{
		opts: opts,
	}
}

type emptyInterface interface {}

func (s *Server) RegisterService(serviceName string, svr interface{}) error {
	svrType := reflect.TypeOf(svr)
	svrValue := reflect.ValueOf(svr)

	sd := &ServiceDesc{
		ServiceName: serviceName,
		Svr: svr,
		HandlerType: (*emptyInterface)(nil),
	}

	methods, err := getServiceMethods(svrType, svrValue)
	if err != nil {
		return err
	}
	sd.Methods = methods

	s.Register(sd, svr)
	return nil
}

func getServiceMethods(serviceType reflect.Type, serviceValue reflect.Value) ([]*MethodDesc, error) {
	var methods []*MethodDesc

	for i := 0; i < serviceType.NumMethod(); i++ {
		method := serviceType.Method(i)
		if err := checkMethod(method.Type); err != nil {
			return nil, err
		}

		methodHandler := func(ctx context.Context, svr interface{}, dec func(interface{}) error) (interface{}, error) {
			reqType := method.Type.In(2)

			req := reflect.New(reqType.Elem()).Interface()

			if err := dec(req); err != nil {
				return nil, err
			}

			handler := func(ctx context.Context, reqBody interface{}) (interface{}, error) {
				values := method.Func.Call([]reflect.Value{serviceValue, reflect.ValueOf(ctx), reflect.ValueOf(req)})
				return values[0].Interface(), nil
			}

			return handler(ctx, req)
		}

		methods = append(methods, &MethodDesc{
			MethodName: method.Name,
			Handler: methodHandler,
		})
	}

	return methods, nil
}

// 检查方法的入参和出参是否合法
func checkMethod(method reflect.Type) error {
	if method.NumIn() < 3 {
		return fmt.Errorf("方法 %s 传入参数不能小于个", method.Name())
	}

	if method.NumOut() != 2 {
		return fmt.Errorf("方法 %s 返回参数必须为 2 个, 实际上返回了 %d 个", method.Name(), method.NumOut())
	}

	ctxType := method.In(1)
	var contextType = reflect.TypeOf((*context.Context)(nil)).Elem()
	if !ctxType.Implements(contextType) {
		return fmt.Errorf("方法 %s 第一个参数不是 context", method.Name())
	}

	argType := method.In(2)
	if argType.Kind() != reflect.Ptr {
		return fmt.Errorf("方法 %s 传入的第二个参数不是一个指针", method.Name())
	}

	replyType := method.Out(0)
	if replyType.Kind() != reflect.Ptr {
		return fmt.Errorf("方法 %s 返回的第一个参数不是一个指针", method.Name())
	}

	errType := method.Out(1)
	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	if !errType.Implements(errorType) {
		return fmt.Errorf("方法 %s 返回的第二个参数不是 error", method.Name())
	}

	return nil
}


func (s *Server) Register(sd *ServiceDesc, svr interface{}) {
	if sd == nil || svr == nil {
		return
	}

	handlerType := reflect.TypeOf(sd.HandlerType).Elem()
	serviceType := reflect.TypeOf(svr)

	if !serviceType.Implements(handlerType) {
		log.Fatalf("handlerType %v not match service: %v", handlerType, serviceType)
	}

	ser := &service{
		svr: svr,
		serviceName: sd.ServiceName,
		handlers: make(map[string]Handler),
	}

	for _, method := range sd.Methods {
		ser.handlers[method.MethodName] = method.Handler
	}

	s.service = ser
}