package server

import (
	mgrApi "austin-v2/api/mgr"
	"austin-v2/app/mgr/internal/conf"
	"austin-v2/app/mgr/internal/middlewarex"
	"austin-v2/app/mgr/internal/service"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	stdHttp "net/http"
)

// NewHTTPServer new an HTTP serviceName.
func NewHTTPServer(c *conf.Server,
	bc *conf.Bootstrap,
	system *service.SystemService,
	common *service.CommonService,
	gmpPlatformService *service.GmpPlatformService,
	logger log.Logger,
) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			// 日志记录
			middlewarex.LoggerServer(logger),
			middlewarex.Jwt(bc.Auth.ApiKey),
			middlewarex.Auth(),
		),
		// 跨域设置
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "AccessToken", "X-Token", "Accept"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	// 增加成功自定义json返回值
	opts = append(opts, http.ResponseEncoder(responseEncoder))
	// 增加错误自定义json返回值
	opts = append(opts, http.ErrorEncoder(errorEncoder))

	srv := http.NewServer(opts...)
	//route := srv.Route("/")
	//route.POST("/files/v1/uploadFile", service.UploadFile)
	mgrApi.RegisterSystemHTTPServer(srv, system)
	mgrApi.RegisterCommonHTTPServer(srv, common)
	mgrApi.RegisterGmpPlatformHTTPServer(srv, gmpPlatformService)
	return srv
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func responseEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, v interface{}) error {
	reply := &Response{}
	reply.Code = 200
	reply.Message = "success"

	codec, _ := http.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(v)
	_ = json.Unmarshal(data, &reply.Data)
	if err != nil {
		return err
	}

	resp, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", codec.Name())
	w.WriteHeader(stdHttp.StatusOK)
	_, _ = w.Write(resp)
	return nil
}

func errorEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, err error) {
	codec, _ := http.CodecForRequest(r, "Accept")
	w.Header().Set("Content-Type", "application/"+codec.Name())
	// 返回码均是200
	w.WriteHeader(stdHttp.StatusOK)
	// 重写errResponse
	err = SetCustomizeErrInfo(err)
	se := errors.FromError(err)
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	_, _ = w.Write(body)
}

// SetCustomizeErrInfo 根据err.Reason返回自定义包装错误
func SetCustomizeErrInfo(err error) error {
	e := errors.FromError(err)
	// 如果 e.Code = 504 则是服务不可达
	if e.Code == stdHttp.StatusGatewayTimeout {
		return errors.New(stdHttp.StatusGatewayTimeout, "SERVICE_GATEWAY_TIMEOUT", "服务不可达")
	}
	// 如果 e.Code = 503 则是服务连接被拒绝
	if e.Code == stdHttp.StatusServiceUnavailable {
		return errors.New(stdHttp.StatusServiceUnavailable, "SERVICE_GATEWAY_UNAVAILABLE", "服务连接被拒绝")
	}
	reason := e.Reason
	msg := e.Message
	code := e.Code
	if reason == "" {
		reason = "UNKNOWN_ERROR"
		msg = "系统异常"
		code = 500
	}
	return errors.New(int(code), reason, msg)
}
