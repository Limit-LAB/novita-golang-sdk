package request

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/novitalabs/golang-sdk/types"
)

func (c *Client) Txt2Img(ctx context.Context, request *types.Txt2ImgRequest) (*types.AsyncResponse, error) {
	responseData, err := doRequest[types.Txt2ImgRequest, types.AsyncResponse](ctx, c.httpCli, http.MethodPost, c.apiPath+"/txt2img", c.apiKey, nil, request)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (c *Client) SyncTxt2img(ctx context.Context, request *types.Txt2ImgRequest, opts ...WithGenerateImageOption) (*types.ProgressResponse, error) {
	return syncImageGeneration[*types.Txt2ImgRequest](ctx, request, opts, c.Txt2Img, c.waitForTask)
}

func (c *Client) waitForTask(ctx context.Context, request *types.ProgressRequest, opts ...WithGenerateImageOption) (*types.ProgressResponse, error) {
	// get sync option
	igOpt := newGenerateImageOption(opts...)
	const checkInterval = time.Second
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			progress, err := c.Progress(ctx, request)
			if err != nil {
				return nil, err
			}
			if !progress.Data.Status.IsFinish() {
				continue
			}
			if progress.Data.Status != types.Successful {
				return nil, fmt.Errorf("generate image failed, %s", progress.Data.FailedReason)
			}
			// other stuff
			if igOpt.DownloadImage {
				if err := progress.Data.DownloadImages(ctx); err != nil {
					return nil, err
				}
			}
			if igOpt.SaveImage {
				if err := os.MkdirAll(igOpt.SaveImageDir, igOpt.SaveImagePerm); err != nil {
					return nil, err
				}
				for i, s3Url := range progress.Data.Imgs {
					lastSlashIndex := strings.LastIndex(s3Url, "/")
					if lastSlashIndex == -1 || lastSlashIndex >= len(s3Url)-1 {
						return nil, fmt.Errorf("can't get file name in url = %s", s3Url)
					}
					fileName := igOpt.SaveImageFileNameConverter(request.TaskId, i, s3Url[lastSlashIndex+1:])
					if err := os.WriteFile(igOpt.SaveImageDir+`/`+fileName, progress.Data.ImgsBytes[i], igOpt.SaveImagePerm); err != nil {
						return nil, err
					}
				}
			}
			return progress, nil
		}
	}
}
