package activitygroups

import (
	"fmt"

	"github.com/mrrizal/todo-api-clean-arch/entity"
)

// func (svc *Service) Get(activityGroupID int) (entity.ActivityGroups, error) {
// 	return entity.ActivityGroups{}, nil
// }

// func (svc *Service) GetAll() ([]entity.ActivityGroups, error) {
// 	return []entity.ActivityGroups{}, nil
// }

func (svc *Service) Create(activityGroup *entity.ActivityGroups) error {
	fmt.Println("haha")
	return nil
}

// func (svc *Service) Update(activityGroup entity.ActivityGroups) (entity.ActivityGroups, error) {
// 	return entity.ActivityGroups{}, nil
// }

// func (svc *Service) Delete(activityGroupID int) error {
// 	return nil
// }
