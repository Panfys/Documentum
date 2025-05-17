package valid

import (
	"github.com/microcosm-cc/bluemonday"
    "documentum/pkg/logger"
)

type validatService struct {
    policy *bluemonday.Policy
    log      logger.Logger
}

func NewValidatService(log logger.Logger) *validatService {
    return &validatService{
        log: log,
        policy: bluemonday.UGCPolicy(),
    }
}
