package web

import (
	"github.com/AVZotov/draft-survey/internal/types"
	"github.com/AVZotov/draft-survey/web/components"
)

func DashboardProps(user *types.User, survey *types.Survey) components.LayoutProps {
	return components.LayoutProps{
		Title:           "Dashboard",
		MetaDescription: "Main page of Draft Survey application",
		ExtraCSS:        []string{"/static/css/dashboard.css"},
		User:            user,
		Survey:          survey,
	}
}

func NewSurveyProps(user *types.User, survey *types.Survey) components.LayoutProps {
	return components.LayoutProps{
		Title:           "New Survey",
		MetaDescription: "Calculate vessel cargo",
		ExtraCSS:        []string{"/static/css/new-survey.css"},
		ExtraJS:         []string{"/static/js/new-survey.js"},
		User:            user,
		Survey:          survey,
	}
}

func DraftReadingsProps(user *types.User, survey *types.Survey) components.LayoutProps {
	return components.LayoutProps{
		Title:           "Drafts Reading",
		MetaDescription: "Get vessel's draft marks",
		ExtraCSS:        []string{"/static/css/draft-readings.css"},
		ExtraJS:         []string{"/static/js/draft-readings.js"},
		User:            user,
		Survey:          survey,
	}
}

var ProfilePageProps = components.LayoutProps{
	Title:           "Surveyor Profile",
	MetaDescription: "Set up your surveyor profile",
}

var BannerFileCorrupted = components.BannerProps{
	Type:    components.Warn,
	Header:  "Profile File Corrupted",
	Message: "Your profile file could not be read. Please fill in your details again.",
}
