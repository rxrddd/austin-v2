package server

import (
	v1 "austin-v2/api/msgpusher/v1"
	"austin-v2/app/msgpusher/internal/conf"
	"austin-v2/app/msgpusher/internal/service"
	"austin-v2/pkg/errResponse"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	stdHttp "net/http"
)

// NewHTTPServer new a HTTP serviceName.
func NewHTTPServer(c *conf.Server, service *service.MsgPusherService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			// 日志记录
			logging.Server(logger),
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
	v1.RegisterMsgPusherHTTPServer(srv, service)
	return srv
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func responseEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, v interface{}) error {
	reply := &Response{}
	reply.Code = 0
	reply.Message = "success"

	codec, _ := http.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(v)
	_ = json.Unmarshal(data, &reply.Data)
	if err != nil {
		return err
	}

	data, err = codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", codec.Name())
	w.WriteHeader(stdHttp.StatusOK)
	w.Write(data)
	return nil
}

func errorEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, err error) {
	codec, _ := http.CodecForRequest(r, "Accept")
	w.Header().Set("Content-Type", "application/"+codec.Name())
	// 返回码均是200
	w.WriteHeader(stdHttp.StatusOK)
	// 重写errResponse
	err = errResponse.SetCustomizeErrInfo(err)
	se := errors.FromError(err)
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	_, _ = w.Write(body)
}
