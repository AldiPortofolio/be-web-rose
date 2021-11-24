package minio

type UploadRequest struct {
	BucketName string
	Data string
	NameFile string
	ContentType string
}

type UploadResponse struct {
	Url string
	NameFile string
	Rc string
	Message string
}