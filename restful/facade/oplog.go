package facade

import "github.com/vela-ssoc/backend-common/model"

type Recorder interface {
	Record(*model.Oplog) error
}

type Describer interface {
	Name() string
	Proc([]byte) []byte
}
