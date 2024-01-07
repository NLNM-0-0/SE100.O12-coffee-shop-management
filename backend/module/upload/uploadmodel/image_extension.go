package uploadmodel

type ImageExtension string

const (
	Png  ImageExtension = ".png"
	Jpeg ImageExtension = ".jpeg"
	Jpg  ImageExtension = ".jpg"
	Icon ImageExtension = ".ico"
)

func CheckImageExtension(extensionFile string) bool {
	extension := ImageExtension(extensionFile)

	switch extension {
	case Png, Jpeg, Jpg, Icon:
		return true
	default:
		return false
	}
}
