package helper

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"resource_service/internal/core/shared"
	"strconv"
	"time"
)

func InitializeLog() {
	logDir := Config.LogDir
	_ = os.Mkdir(logDir, os.ModePerm)

	f, err := os.OpenFile(logDir+Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetFlags(0)
	log.SetOutput(f)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func LogRequest(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	statusCode := c.Writer.Status()
	response := shared.NoErrorsFound
	level := "INFO"
	if statusCode >= 400 {
		response = blw.body.String()
		level = "ERROR"
	}
	data, err := json.Marshal(&LogStruct{
		Method:          c.Request.Method,
		Level:           level,
		StatusCode:      strconv.Itoa(statusCode),
		Path:            c.Request.URL.String(),
		UserAgent:       c.Request.Header.Get("User-Agent"),
		RemoteIP:        c.ClientIP(),
		ResponseTime:    time.Since(time.Now()).String(),
		Message:         http.StatusText(statusCode) + " : " + response,
		Version:         "1",
		CorrelationId:   uuid.New().String(),
		AppName:         Config.AppName,
		ApplicationHost: c.Request.Host,
		LoggerName:      "",
		TimeStamp:       time.Now().Format(time.RFC3339),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", data)
	c.Next()

}
func LogEvent(level string, message interface{}) {

	data, err := json.Marshal(struct {
		TimeStamp string      `json:"@timestamp"`
		Level     string      `json:"level"`
		Message   interface{} `json:"message"`
	}{TimeStamp: time.Now().Format(time.RFC3339),
		Message: message,
		Level:   level,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", data)

}
