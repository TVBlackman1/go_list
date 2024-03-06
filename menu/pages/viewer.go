package pages

type PageViewer interface {
	View()
	Options() PageViewer
}
