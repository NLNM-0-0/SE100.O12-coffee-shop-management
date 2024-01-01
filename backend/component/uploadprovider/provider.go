package uploadprovider

import "backend/common"

type UploadStaticProvider interface {
	UploadImage(data []byte, path string) (common.Image, error)
}
