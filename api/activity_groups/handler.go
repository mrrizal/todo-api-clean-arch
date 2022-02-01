package activitygroups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/todo-api-clean-arch/api/utils"
	"github.com/mrrizal/todo-api-clean-arch/entity"
	activitygroups_usecase "github.com/mrrizal/todo-api-clean-arch/usecase/activity_groups"
)

var ActivityGroupsUsecase activitygroups_usecase.UseCase

func CreateActivityGroupsHandler(c *fiber.Ctx) error {
	var activityGroups entity.ActivityGroups
	if err := c.BodyParser(&activityGroups); err != nil {
		return utils.Response(c, 400, utils.ErrorResponse(err.Error()))
	}

	if err := activityGroups.Validate(); err != nil {
		return utils.Response(c, 400, utils.ErrorResponse(err.Error()))
	}

	ActivityGroupsUsecase.Create(&entity.ActivityGroups{})
	c.Status(201)
	return nil
}
