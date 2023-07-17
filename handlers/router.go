package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"

	"widget-sports/configurations"
	"widget-sports/services"
)

const message = "некорректный URL"

func Handlers(cfg configurations.Config) {

	app := fiber.New()
	app.Get("/widget-sport/:sport/:endpoint", func(c *fiber.Ctx) error {
		return GetMatches(c, cfg)
	})

	err := app.Listen(":" + cfg.Port)
	if err != nil {
		log.Panic(err)
	}
}

func GetMatches(c *fiber.Ctx, cfg configurations.Config) error {
	sport := c.Params("sport")
	endpoint := c.Params("endpoint")

	var response interface{}
	var err error

	switch endpoint {
	case "fixtures-by-date":
		switch sport {
		case "basketball":
			response, err = services.GetBasketballMatchesByDate(cfg)
		case "football":
			response, err = services.GetFootballMatchesByDate(cfg)
		case "hockey":
			response, err = services.GetHockeyMatchesByDate(cfg)
		case "mma":
			response, err = services.GetMMAMatchesByDate(cfg)
		case "tennis":
			response, err = services.GetTennisMatchesByDate(cfg)
		default:
			err = errors.New(message)
		}
	case "fixtures-in-progress":
		switch sport {
		case "basketball":
			response, err = services.GetBasketballMatchesByLiveScore(cfg)
		case "football":
			response, err = services.GetFootballMatchesByLiveScore(cfg)
		case "hockey":
			response, err = services.GetHockeyMatchesByLiveScore(cfg)
		case "tennis":
			response, err = services.GetTennisMatchesByLiveScore(cfg)
		default:
			err = errors.New(message)
		}
	case "results":
		switch sport {
		case "mma":
			response, err = services.GetMMAMatchesResults(cfg)
		default:
			err = errors.New(message)
		}
	default:
		err = errors.New(message)
	}

	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("error marshaling response:", err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	c.Set("Content-Type", "application/json")
	return c.Status(http.StatusOK).Send(jsonResponse)
}
