package todo

import "github.com/mrrizal/todo-api-clean-arch/entity"

type Repository interface {
	Get(todoID int) (entity.Todo, error)
	GetAll(activityGroupID int) ([]entity.Todo, error)
	Create(todo entity.Todo) (entity.Todo, error)
	Update(todo entity.Todo) (entity.Todo, error)
	Delete(todoID int) error
}
