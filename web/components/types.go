package components

import (
	"github.com/AVZotov/draft-survey/internal/calculation"
	"github.com/AVZotov/draft-survey/internal/types"
)

type MetaData struct {
	SurveyID      string
	DraftIndex    string
	DratType      string
	TotalBwWeight string
	TotalFwWeight string
}

type BwTankCorrections struct {
	Tank *types.BallastWaterTank
	Trim *float64
	List *float64
}

type LayoutProps struct {
	Title           string
	MetaDescription string
	User            *types.User
	Survey          *types.Survey
	Results         *[]calculation.DraftResult
	ExtraCSS        []string
	ExtraJS         []string
	MetaData        *MetaData
	*BwTankCorrections
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

type ModalLevel string

const (
	ModalInfo    ModalLevel = "Information"
	ModalWarning ModalLevel = "Warning"
	ModalError   ModalLevel = "Danger"
)

type ModalProps struct {
	Level      ModalLevel
	DialogID   string
	Title      string
	Message    string
	ConfirmBtn string
	CancelBtn  string // Do not rendered if prop is empty
}
