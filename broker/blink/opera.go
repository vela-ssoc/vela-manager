package blink

import "net/http"

var BrkPing = &opera{method: http.MethodGet, path: "/ping"}

type Operator interface {
	Method() string
	Path() string
	Desc() string
}

type opera struct {
	method string
	path   string
	desc   string
}

func (op *opera) Method() string { return op.method }
func (op *opera) Path() string   { return op.path }
func (op *opera) Desc() string   { return op.desc }
