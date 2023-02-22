package blink

import "net/http"

var (
	BrkPing = &opera{method: http.MethodGet, path: "/api/ping"}
	BrkEnv  = &opera{method: http.MethodGet, path: "/api/env"}
)

type Operator interface {
	Method() string
	Path() string
	Desc() string
}

func NewOp(method, path, desc string) Operator {
	return &opera{
		method: method,
		path:   path,
		desc:   desc,
	}
}

type opera struct {
	method string
	path   string
	desc   string
}

func (op *opera) Method() string { return op.method }
func (op *opera) Path() string   { return op.path }
func (op *opera) Desc() string   { return op.desc }
