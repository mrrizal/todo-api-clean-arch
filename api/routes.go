package api

import (
	"github.com/gofiber/fiber/v2"
	activitygroups "github.com/mrrizal/todo-api-clean-arch/api/activity_groups"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/activity-groups/", activitygroups.CreateActivityGroupsHandler)
}
