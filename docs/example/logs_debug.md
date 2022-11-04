## 记录请求详细日志
```go
package middleware

import (
	"app"
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

var debugLog *os.File

func init() {

	debugLog, _ = os.OpenFile(app.Dir.Data+"debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}

func LoggingDebugMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var reader io.Reader = r.Body
		b, _ := io.ReadAll(reader)
		r.Body = io.NopCloser(bytes.NewReader(b)) // reuse body

		if debugLog != nil {
			head := ""
			for k, v := range r.Header {
				head += k + ": " + v[0] + "\n"
			}
			debugLog.WriteString("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
			debugLog.WriteString(time.Now().String() + "\n")
			debugLog.WriteString(r.Method + " " + r.URL.String() + "\n")
			debugLog.WriteString(head + "\n")
			if len(b) > 1024*100 {
				debugLog.WriteString("[Too Large Size]")
				debugLog.WriteString("\n")
			} else if len(b) > 0 {
				debugLog.Write(b)
				debugLog.WriteString("\n")
			}
		}

		next.ServeHTTP(w, r)
	})
}

```
