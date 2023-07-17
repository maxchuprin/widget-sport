package services

import (
	"fmt"
	"strings"
	"time"
	"widget-sports/configurations"
	"widget-sports/request"
	"widget-sports/structures"
	"widget-sports/utils"
)

type respBasketball struct {
	Response []structures.BasketballMatch `json:"response"`
}

func GetBasketballMatchesByDate(cfg configurations.Config) ([]map[string]interface{}, error) {
	cacheKey := "basketballByDate"
	matches, found := utils.GetFromCache(cacheKey)
	if found {
		return matches, nil
	}

	url := fmt.Sprintf("https://%s/games?timezone=%s&date=%s", cfg.BasketballHost, utils.CheckTimeZone(cfg.TimeZone), time.Now().AddDate(0, 0, cfg.PlusDay).Format("2006-01-02"))

	var r respBasketball
	bodyResponse, err := request.ServerRequest(cfg.BasketballHost, cfg.RapidAPIKey, url, &r)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	err = utils.GetMessage(bodyResponse)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	var output []map[string]interface{}
	for _, v := range r.Response {
		if utils.Contains(strings.Split(cfg.BasketballLeague, ","), v.League.Name) {
			item := map[string]interface{}{
				"matchDateTime": v.Date,
				"leagueName":    v.League.Name,
				"leagueLogo":    v.League.Logo,
				"homeTeamName":  v.Teams.Home.Name,
				"homeTeamLogo":  v.Teams.Home.Logo,
				"awayTeamName":  v.Teams.Away.Name,
				"awayTeamLogo":  v.Teams.Away.Logo,
			}
			output = append(output, item)
		}
	}
	utils.SetCache(cacheKey, output)

	return output, nil
}

func GetBasketballMatchesByLiveScore(cfg configurations.Config) ([]map[string]interface{}, error) {
	cacheKey := "basketballLiveScore"
	matches, found := utils.GetFromCache(cacheKey)
	if found {
		return matches, nil
	}

	url := fmt.Sprintf("https://%s/games?timezone=%s&date=%s", cfg.BasketballHost, cfg.TimeZone, time.Now().Format("2006-01-02"))

	var r respBasketball

	bodyResponse, err := request.ServerRequest(cfg.BasketballHost, cfg.RapidAPIKey, url, &r)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	err = utils.GetMessage(bodyResponse)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	var output []map[string]interface{}
	for _, v := range r.Response {
		if utils.Contains(strings.Split(cfg.BasketballLeague, ","), v.League.Name) {
			switch v.Status.Period {
			case "Q1", "Q2", "Q3", "Q4", "OT", "BT", "HT":
				item := map[string]interface{}{
					"timer":        v.Status.Timer,
					"period":       v.Status.Period,
					"leagueName":   v.League.Name,
					"leagueLogo":   v.League.Logo,
					"homeTeamName": v.Teams.Home.Name,
					"homeTeamLogo": v.Teams.Home.Logo,
					"homeScore":    v.Scores.Home,
					"awayTeamName": v.Teams.Away.Name,
					"awayTeamLogo": v.Teams.Away.Logo,
					"awayScore":    v.Scores.Away,
				}
				output = append(output, item)
			}
		}
	}
	utils.SetCache(cacheKey, output)

	return output, nil
}
