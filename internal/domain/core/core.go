package core

import (
	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/adapters/logger"
)

type St struct {
	lg    logger.Lite
	httpc httpc.HttpC
}

func New(lg logger.Lite, httpc httpc.HttpC) *St {
	return &St{
		lg:    lg,
		httpc: httpc,
	}
}
