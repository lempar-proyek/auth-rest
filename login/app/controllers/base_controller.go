package controllers

import (
	"net/http"

	"github.com/lempar-proyek/login-rest/login/app/dto"
	"github.com/revel/revel"
)

type BaseController struct {
	*revel.Controller
}

func (c *BaseController) EndPointNotFound() revel.Result {
	c.Response.Status = http.StatusNotFound

	response := dto.ErrorResponse{}
	response.Code = http.StatusNotFound
	response.Error = "EP_NOT_FOUND"
	response.Messages = "end point not found"

	return c.RenderJSON(response)
}
