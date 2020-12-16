package handler

import (
	"fmt"
	"github.com/B8kingS0ga/Go-000/tree/main/Week04/internal/api/service"
	"net/http"
)

type HelloHandle struct {
	Service service.Service
}

func (h *HelloHandle) Start() {
	http.Handle("/hello", h)
	http.ListenAndServe(":8090", nil)
}

func (h *HelloHandle) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "greeting: "+h.Service.ServiceSome())
}

func NewHandle(s service.Service) HelloHandle {
	return HelloHandle{
		Service: s,
	}
}
