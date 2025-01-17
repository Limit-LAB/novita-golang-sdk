package request

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/novitalabs/golang-sdk/types"
	"github.com/novitalabs/golang-sdk/util"
)

func TestClient_SyncImg2Img(t *testing.T) {
	client, err := NewClient(os.Getenv("API_KEY"), "")
	if err != nil {
		t.Error(err)
		return
	}
	initImage := "out/test_txt2img_sync.png"
	initImageBase64, err := util.ReadImageToBase64(initImage)
	if err != nil {
		t.Error(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	img2Img := types.NewImg2ImgRequest("a dog flying in the sky", "", "AnythingV5_v5PrtRE.safetensors", initImageBase64)
	res, err := client.SyncImg2img(ctx, img2Img,
		WithSaveImage("out", 0777, func(taskId string, fileIndex int, fileName string) string {
			return "test_img2img_sync.png"
		}))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("status = %d", res.Data.Status)
}

func TestClient_SyncImg2imgControlNet(t *testing.T) {
	client, err := NewClient(os.Getenv("API_KEY"), "")
	if err != nil {
		t.Error(err)
		return
	}
	initImage := "out/test_img2img_sync.png"
	initImageBase64, err := util.ReadImageToBase64(initImage)
	if err != nil {
		t.Error(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	img2Img := types.NewImg2ImgRequest("a dog flying in the sky", "", "AnythingV5_v5PrtRE.safetensors", initImageBase64)
	controlNetReq := types.NewControlNetUnit(types.Canny, "control_v11p_sd15_canny", initImageBase64)
	img2Img.ControlNetUnits = []*types.ControlNetUnit{controlNetReq}
	res, err := client.SyncImg2img(ctx, img2Img,
		WithSaveImage("out", 0777, func(taskId string, fileIndex int, fileName string) string {
			if fileIndex == 0 {
				return "test_img2img_controlnet_sync.png"
			} else {
				return "test_img2img_controlnet_processor.png"
			}
		}))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("status = %d", res.Data.Status)
}
