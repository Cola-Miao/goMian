package log

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"
)

func Init() error {
	fp, err := makeLogPath("logs")
	if err != nil {
		return err
	}
	opt := slog.HandlerOptions{
		AddSource:   false,
		Level:       nil,
		ReplaceAttr: nil,
	}
	h := slog.NewJSONHandler(fp, &opt)
	l := slog.New(h)
	slog.SetDefault(l)
	return nil
}

func makeLogPath(path string) (*os.File, error) {
	date := time.Now().Format("/200601")
	if err := os.MkdirAll(path+date, os.ModePerm); err != nil {
		return nil, err
	}
	file := fmt.Sprintf("/%s.log", strconv.Itoa(time.Now().Day()))
	fp, err := os.OpenFile(path+date+file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	return fp, err
}
