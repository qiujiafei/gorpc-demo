package gorpc_demo

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


	return nil
}


func (s *Server) Register(svr interface{}) {

}