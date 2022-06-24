package core

import (
	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/adapters/logger"
)

type St struct {
	lg         logger.Lite
	httpc      httpc.HttpC
	senderName string
}

func New(lg logger.Lite, httpc httpc.HttpC, senderName string) *St {
	return &St{
		lg:         lg,
		httpc:      httpc,
		senderName: senderName,
	}
}
