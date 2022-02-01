package activitygroups

import "github.com/mrrizal/todo-api-clean-arch/entity"

type Repository interface {
	Get(activityGroupID int) (entity.ActivityGroups, error)
	GetAll() ([]entity.ActivityGroups, error)
	Create(activityGroup entity.ActivityGroups) (entity.ActivityGroups, error)
	Update(activityGroup entity.ActivityGroups) (entity.ActivityGroups, error)
	Delete(activityGroupID int) error
}
