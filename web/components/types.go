package components

type LayoutProps struct {
	Title           string
	MetaDescription string
}

type Version struct {
	Version string
}

type BannerType string

const (
	Info  BannerType = "info"
	Warn  BannerType = "warn"
	Error BannerType = "error"
)

type BannerProps struct {
	Type    BannerType
	Header  string
	Message string
	Details string
}
