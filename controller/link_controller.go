package controller

import (
	"ShortLink/service"
	"net/http"
	"strings"
)

type LinkController struct {
	service *service.LinkService
}

func NewLinkController(service *service.LinkService) *LinkController {
	return &LinkController{service: service}
}

func (c *LinkController) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.TrimPrefix(r.URL.Path, "/")

	originalURL, err := c.service.GetOriginalURL(shortCode)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}
