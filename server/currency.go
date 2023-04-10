package server

import (
	"context"
	"gprc/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	logger hclog.Logger
}

func (c *Currency) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {
	c.logger.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &currency.RateResponse{Rate: 0.5}, nil
}

func (c *Currency) mustEmbedUnimplementedCurrencyServer() {
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}
