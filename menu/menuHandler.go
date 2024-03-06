package menu

import (
	"tt/menu/pages"

	"github.com/inancgumus/screen"
)

type MenuHandler struct {
	currentPage pages.PageViewer
}

func NewMenuHandler() *MenuHandler {
	return &MenuHandler{
		currentPage: &pages.Welcome{},
	}
}

func (h *MenuHandler) Run() {
	for {
		screen.Clear()
		screen.MoveTopLeft()

		h.currentPage.View()
		h.currentPage = h.currentPage.Options()
	}
}
