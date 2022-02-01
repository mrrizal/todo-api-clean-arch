package activitygroups

import "github.com/mrrizal/todo-api-clean-arch/entity"

func Get(activitygGroupID int) (entity.ActivityGroups, error) {
	return entity.ActivityGroups{}, nil
}

func GetAll() ([]entity.ActivityGroups, error) {
	return []entity.ActivityGroups{}, nil
}

func Create(activityGroup entity.ActivityGroups) (entity.ActivityGroups, error) {
	return entity.ActivityGroups{}, nil
}

func Update(activityGroup entity.ActivityGroups) (entity.ActivityGroups, error) {
	return entity.ActivityGroups{}, nil
}

func Delete(activityGroupID int) error {
	return nil
}
