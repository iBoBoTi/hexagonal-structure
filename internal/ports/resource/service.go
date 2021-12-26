package port

import domain "resource_service/internal/core/domain/resource"

type ResourceService interface {
	Create(resource domain.Resource) (interface{}, error)
	Read(reference string) (interface{}, error)
	Update(reference string, resource domain.Resource) (interface{}, error)
	Delete(reference string) (interface{}, error)
	ReadAll() (interface{}, error)
}
