package services

import (
	domain "resource_service/internal/core/domain/resource"
	"resource_service/internal/core/helper"
	port "resource_service/internal/ports/resource"
)

type service struct {
	resourceRepository port.ResourceRepository
}

func (service service) Update(reference string, resource domain.Resource) (interface{}, error) {
	_, err := service.Read(reference)
	if err != nil {
		return nil, err
	}
	if err := helper.Validate(resource); err != nil {
		return nil, err
	}
	return service.resourceRepository.Update(reference, resource)
}
func (service service) Delete(reference string) (interface{}, error) {
	_, err := service.Read(reference)
	if err != nil {
		return nil, err
	}
	return service.resourceRepository.Delete(reference)
}
func New(resourceRepository port.ResourceRepository) *service {
	return &service{
		resourceRepository: resourceRepository,
	}
}
func (service *service) Read(reference string) (interface{}, error) {
	resource, err := service.resourceRepository.Read(reference)
	if err != nil {
		return nil, err
	}
	return resource, nil
}
func (service *service) ReadAll() (interface{}, error) {
	resources, err := service.resourceRepository.ReadAll()
	if err != nil {
		return nil, err
	}
	return resources, nil
}
func (service *service) Create(resource domain.Resource) (interface{}, error) {
	if err := helper.Validate(resource); err != nil {
		return nil, err
	}
	return service.resourceRepository.Create(resource)
}
