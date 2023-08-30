package middlewares

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"mumu/pkg/logger"
	"mumu/pkg/view"
	"net/http"
	"net/http/httputil"
	"time"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}

				stackTrace := zap.Stack("stacktrace")
				httpRequest, _ := httputil.DumpRequest(request, true)
				logger.Error("recovery from panic",
					zap.Time("time", time.Now()),               // 记录时间
					zap.Any("error", err),                      // 记录错误信息
					zap.String("request", string(httpRequest)), // 请求信息
					stackTrace,                                // 调用堆栈信息
				)

				w.WriteHeader(500)
				view.Render(w, view.D{"msg": "服务器发生了异常：" + err.Error()}, "errors/panic")
				sendMeMail(err, stackTrace)
			}
		}()
		next.ServeHTTP(w, request)
	})
}

func sendMeMail(err error, stack zapcore.Field) {
	// send mail,编写逻辑，把错误发送到邮箱或钉钉
	//fmt.Println("邮件通知，这里出现了一个panic:" + err.Error() + stack.String)
}
