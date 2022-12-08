package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	v1 "pay/api/pay/v1"
	"pay/internal/conf"
	"pay/internal/service"
	"strings"
)

func notifyMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		var op string
		if info, ok := transport.FromServerContext(ctx); ok {
			op = info.Operation()
		}
		log.Log(log.LevelDebug, "op", op)
		if !strings.Contains(op, "notify") {
			return handler(ctx, req)
		}
		reply, err = handler(ctx, req)
		if err != nil {
			return "success", nil
		}
		return
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, pay *service.PayService) *http.Server {
	opts := []http.ServerOption{
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "HEAD"}),
			handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Content-Type"}),
		)),
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			validate.Validator(),
			notifyMiddleware,
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
	srv := http.NewServer(opts...)
	v1.RegisterPayServiceHTTPServer(srv, pay)
	return srv
}
