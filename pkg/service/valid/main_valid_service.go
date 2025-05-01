package valid

import (
	"github.com/microcosm-cc/bluemonday"
)

type validatService struct {
    policy *bluemonday.Policy
}

func NewValidatService() *validatService {
    return &validatService{
        policy: bluemonday.UGCPolicy(),
    }
}
