package controllers

import (
	"github.com/acme-corp/app/models"
	"github.com/acme-corp/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rs/xid"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
)

func AddTeam(c *fiber.Ctx) error {
	team := &models.Team{}
	// Check, if received JSON data is valid.
	if err := c.BodyParser(team); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	team.ID = xid.New().String()
	if err := database.DB.Create(&team).Error; err != nil {
		log.Info("Creating team failed: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	roles := []string{"owner", "admin", "member", "customer"}
	for _, role := range roles {
		log.Info("Creating roles")
		res, err := userroles.CreateNewRoleOrAddPermissions(team.ID+"_"+role, []string{
			"read",
		}, nil)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		log.Info(res.OK.CreatedNewRole)
	}
	userID := c.GetReqHeaders()["X-User"][0]
	err := database.DB.Model(&team).Association("Users").Append(&models.User{ID: userID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	response, err := userroles.AddRoleToUser("public", userID, team.ID+"_owner", nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if response.UnknownRoleError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   response.UnknownRoleError,
		})

	}

	sessionContainer := session.GetSessionFromRequestContext(c.Context())
	accessTokenPayload := sessionContainer.GetAccessTokenPayload()
	// var teams []models.AccessTokenTeamPayload
	teams := accessTokenPayload["teams"].([]interface{})
	newTeam := map[string]interface{}{"id": team.ID, "name": team.Name}
	teams = append(teams, newTeam)
	accessTokenPayload["teams"] = teams
	if err := sessionContainer.MergeIntoAccessTokenPayload(accessTokenPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  team,
	})
}
func GetTeams(c *fiber.Ctx) error {
	log.Info(c.GetReqHeaders()["X-User"])
	teams := []models.Team{}
	database.DB.Preload("Users").Find(&teams)
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  teams,
	})
}
