package api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	domain "resource_service/internal/core/domain/resource"
	"resource_service/internal/core/helper"
	"resource_service/internal/core/shared"
	port "resource_service/internal/ports/resource"
)

type HTTPHandler struct {
	resourceService port.ResourceService
}

func NewHTTPHandler(resourceService port.ResourceService) *HTTPHandler {
	return &HTTPHandler{
		resourceService: resourceService,
	}
}
func (hdl *HTTPHandler) Read(c *gin.Context) {
	resource, err := hdl.resourceService.Read(c.Param("reference"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}
	if resource == nil {
		_ = c.AbortWithError(404, helper.PrintErrorMessage("404", shared.NoRecordFound))
		return
	}
	c.JSON(200, resource)
}
func (hdl *HTTPHandler) ReadAll(c *gin.Context) {
	resources, err := hdl.resourceService.ReadAll()

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}
	if reflect.ValueOf(resources).IsNil() {
		c.AbortWithStatusJSON(404, helper.PrintErrorMessage("404", shared.NoRecordFound))
		return
	}

	c.JSON(200, resources)
}
func (hdl *HTTPHandler) Create(c *gin.Context) {
	body := domain.Resource{}
	_ = c.BindJSON(&body)

	resource, err := hdl.resourceService.Create(body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(200, gin.H{"reference": resource})
}
func (hdl *HTTPHandler) Update(c *gin.Context) {
	body := domain.Resource{}
	_ = c.BindJSON(&body)
	resource, err := hdl.resourceService.Update(c.Param("reference"), body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, gin.H{"reference": resource})
}
func (hdl *HTTPHandler) Delete(c *gin.Context) {

	resource, err := hdl.resourceService.Delete(c.Param("reference"))
	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}

	c.JSON(200, gin.H{"reference": resource})
}
