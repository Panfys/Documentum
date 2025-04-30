package valid

import (
	"github.com/microcosm-cc/bluemonday"
)

type Validator struct {
    policy *bluemonday.Policy
}

func NewValidator() *Validator {
    return &Validator{
        policy: bluemonday.UGCPolicy(),
    }
}
