package config

import (
	"fmt"
	"os"
	"time"
	"runtime"
	"io"
	"net/http"

)

// Level
const (
	Critical
	Error
	Warning 
	Good
	Info
) 

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	grey    = "\033[97;48m"
	reset   = "\033[0m"
)

func color(color, int) string {
	return fmt.Sprintf("\033[%dm", int(color))	
}


type logForm struct {
	Request *http.Request

	TimeStamp time.Time

	StatusCode int

	Latency time.Duration

	keys map[string] interface{}
}

func (b *logForm) colorForm() string {
	code = b.StatusCode

}

func (b *colorForm) MethodColor() string {
	method := b.method

	switch method {
	case http.MethodGet:
		return grey
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}

}


func initcolors(
	colors = map[Level]string{
		Critical:
	}
)