package validate

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// TranError 翻译后的错误信息
type TranError struct {
	trans validator.ValidationErrorsTranslations
}

func (e *TranError) Error() string {
	causes := make([]string, 0, len(e.trans))
	for _, cause := range e.trans {
		causes = append(causes, cause)
	}

	return strings.Join(causes, ", ")
}
