package components

import (
	"github.com/AVZotov/draft-survey/internal/calculation"
	"github.com/AVZotov/draft-survey/internal/types"
)

type LayoutProps struct {
	Title           string
	MetaDescription string
	User            *types.User
	Survey          *types.Survey
	Results         *[]calculation.DraftResult
	ExtraCSS        []string
	ExtraJS         []string
}

type Version struct {
	Version string
}

type DraftBlockProps struct {
	Draft  types.Draft
	Index  int
	Prefix string
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
