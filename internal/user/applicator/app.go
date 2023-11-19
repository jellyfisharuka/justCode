package applicator

import (
	"callboard/internal/auth/config"
	

	"go.uber.org/zap"
)
type Applicator struct {
	logger *zap.SugaredLogger
	config *config.Config
}

func NewApplicator(logger *zap.SugaredLogger, config *config.Config) *Applicator {
	return &Applicator{
		logger: logger,
		config: config,
	}
}
