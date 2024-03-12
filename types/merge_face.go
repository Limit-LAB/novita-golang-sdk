package types

type MergeFaceRequest struct {
	FaceImageFile string                `json:"face_image_file"`
	ImageFile     string                `json:"image_file"`
	Extra         MergeFaceRequestExtra `json:"extra"`
}

type MergeFaceRequestExtra struct {
	ResponseImageType string `json:"response_image_type"`
	EnterprisePlan    struct {
		Enable bool `json:"enable"`
	} `json:"enterprise_plan"`
}

func NewMergeFaceRequest(faceImageFile, imageFile string) *MergeFaceRequest {
	return &MergeFaceRequest{
		FaceImageFile: faceImageFile,
		ImageFile:     imageFile,
		Extra: MergeFaceRequestExtra{
			ResponseImageType: "jpeg",
			EnterprisePlan: struct {
				Enable bool `json:"enable"`
			}{
				Enable: false,
			},
		},
	}
}

type MergeFaceResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`

	ImageFile string `json:"image_file"`
	ImageType string `json:"image_type"`
}

func (r MergeFaceResponse) GetCode() int {
	return r.Code
}

func (r MergeFaceResponse) GetMsg() string {
	return r.Msg
}
