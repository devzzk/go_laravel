package controllers

import "github.com/devzzk/go_laravel/src/services"

type Controller struct {
	Service *services.Service
}

func (c Controller) SetService(service *services.Service) {
	c.Service = service
}
