package web

import "github.com/AVZotov/draft-survey/web/components"

var HomePageProps = components.LayoutProps{
	Title:           "Draft Survey calculator",
	MetaDescription: "Application to measure cargo weight loaded or discharged from a vessel",
}

var ProfilePageProps = components.LayoutProps{
	Title:           "Surveyor Profile",
	MetaDescription: "Set up your surveyor profile",
}

var DashboardPageProps = components.LayoutProps{
	Title:           "Dashboard",
	MetaDescription: "Main page of Draft Survey application",
	ExtraCSS:        []string{"/static/css/dashboard.css"},
}

var BannerFileCorrupted = components.BannerProps{
	Type:    components.Warn,
	Header:  "Profile File Corrupted",
	Message: "Your profile file could not be read. Please fill in your details again.",
}
