package request

import (
	"context"
	"net/http"

	"github.com/novitalabs/golang-sdk/types"
)

func (c *Client) Models(ctx context.Context, opts ...WithModelOption) (types.ModelList, error) {
	modelOpt := newModelOption(opts...)
	if c.modelCache == nil || len(c.modelCache) == 0 || modelOpt.Refresh {
		responseData, err := doRequest[*types.ModelRequest, types.ModelsResponse](ctx, c.httpCli, http.MethodGet, c.apiPath+"/models", c.apiKey, nil, nil)
		if err != nil {
			return nil, err
		}
		c.modelCache = responseData.Data.Models
	}
	return c.modelCache, nil
}

type ModelOption struct {
	Refresh bool
}

func newModelOption(opts ...WithModelOption) *ModelOption {
	all := &ModelOption{}
	for _, opt := range opts {
		opt(all)
	}
	return all
}

type WithModelOption func(opt *ModelOption)

func WithRefresh() WithModelOption {
	return func(opt *ModelOption) {
		opt.Refresh = true
	}
}
