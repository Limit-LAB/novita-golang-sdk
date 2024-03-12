package request

import (
	"context"
	"net/http"

	"github.com/novitalabs/golang-sdk/types"
)

func (c *Client) MergeFace(ctx context.Context, request *types.MergeFaceRequest) (*types.MergeFaceResponse, error) {
	responseData, err := doRequest[types.MergeFaceRequest, types.MergeFaceResponse](
		ctx, c.httpCli, http.MethodPost, c.apiPath+"/v3/merge-face", c.apiKey, nil, request,
	)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
