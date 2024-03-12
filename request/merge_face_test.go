package request

import (
	"context"
	"encoding/base64"
	"io"
	"os"
	"testing"

	"github.com/novitalabs/golang-sdk/types"
)

func TestClient_MergeFace(t *testing.T) {
	client, err := NewClient(os.Getenv("API_KEY"), "")
	if err != nil {
		t.Error(err)
		return
	}

	image, err := readFileToBase64(os.Getenv("IMAGE_FILE"))
	if err != nil {
		t.Error(err)
		return
	}
	faceImage, err := readFileToBase64(os.Getenv("FACE_IMAGE_FILE"))
	if err != nil {
		t.Error(err)
		return
	}

	response, err := client.MergeFace(context.Background(), &types.MergeFaceRequest{
		ImageFile:     string(image),
		FaceImageFile: string(faceImage),
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(response)
	err = writeBase64ToFile(os.Getenv("OUTPUT_FILE"), []byte(response.ImageFile))
	if err != nil {
		t.Error(err)
		return
	}
}

func readFileToBase64(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	raw, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, base64.StdEncoding.EncodedLen(len(raw)))
	base64.StdEncoding.Encode(buf, raw)

	return buf, nil
}

func writeBase64ToFile(path string, raw []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	data := make([]byte, base64.StdEncoding.DecodedLen(len(raw)))
	_, err = base64.StdEncoding.Decode(data, raw)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
