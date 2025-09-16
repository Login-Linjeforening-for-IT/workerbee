package handlers

import "gitlab.login.no/tekkom/web/beehive/admin-api/v2/services"

type Handler struct {
	Events services.EventService
	Stats  services.StatsService
	Forms  services.FormService
}