package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"

	//"log"
	"net/http"
	"strings"
)

// 限制显示body内容最大长度
const limitSize = 300

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func getBodyData(buf *bytes.Buffer) string {
	var body string
	if buf.Len() > limitSize {
		body = string(buf.Bytes()[:limitSize]) + " ...... "
		// 如果有敏感数据需要过滤掉，比如明文密码
	} else {
		body = buf.String()
	}
	return body
}

func LogoInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		//start := time.Now()
		//  处理前打印输入信息
		buf := bytes.Buffer{}
		_, _ = buf.ReadFrom(c.Request.Body)
		if c.Request.Method == http.MethodPost ||
			c.Request.Method == http.MethodPut ||
			c.Request.Method == http.MethodPatch {
			log.Println("request body -> ", getBodyData(&buf), "\n")
		}

		c.Request.Body = ioutil.NopCloser(&buf)

		//  替换writer
		newWriter := &bodyLogWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = newWriter

		//  处理请求
		c.Next()

		// 处理后打印返回信息
		log.Println("response body -> ", strings.TrimRight(getBodyData(newWriter.body), "\n"))
	}
}
