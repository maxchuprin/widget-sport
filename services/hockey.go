package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"widget-sports/configurations"
	"widget-sports/request"
	"widget-sports/structures"
	"widget-sports/utils"
)

type respHockey struct {
	Response []structures.HockeyMatch `json:"response"`
}

func GetHockeyMatchesByDate(cfg configurations.Config) ([]map[string]interface{}, error) {
	cacheKey := "hockeyByDate"
	matches, found := utils.GetFromCache(cacheKey)
	if found {
		return matches, nil
	}

	url := fmt.Sprintf("https://%s/games/?date=%s&timezone=%s", cfg.HockeyHost, time.Now().AddDate(0, 0, cfg.PlusDay).Format("2006-01-02"), utils.CheckTimeZone(cfg.TimeZone))

	var r respHockey

	bodyResponse, err := request.ServerRequest(cfg.HockeyHost, cfg.RapidAPIKey, url, &r)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	err = utils.GetMessage(bodyResponse)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	var output []map[string]interface{}
	for _, v := range r.Response {
		if utils.Contains(strings.Split(cfg.HockeyLeague, ","), v.League.Name) {
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

func GetHockeyMatchesByLiveScore(cfg configurations.Config) ([]map[string]interface{}, error) {
	cacheKey := "hockeyLiveScore"
	matches, found := utils.GetFromCache(cacheKey)
	if found {
		return matches, nil
	}

	url := fmt.Sprintf("https://%s/games/?date=%s&timezone=%s", cfg.HockeyHost, time.Now().Format("2006-01-02"), utils.CheckTimeZone(cfg.TimeZone))

	var r respHockey

	bodyResponse, err := request.ServerRequest(cfg.HockeyHost, cfg.RapidAPIKey, url, &r)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	err = utils.GetMessage(bodyResponse)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	var output []map[string]interface{}
	for _, v := range r.Response {
		if utils.Contains(strings.Split(cfg.HockeyLeague, ","), v.League.Name) {
			switch v.Status.Period {
			case "P1", "P2", "P3", "OT", "PT", "BT":
				item := map[string]interface{}{
					"leagueName":   v.League.Name,
					"leagueLogo":   v.League.Logo,
					"homeTeamName": v.Teams.Home.Name,
					"homeTeamLogo": v.Teams.Home.Logo,
					"homeScore":    strconv.Itoa(v.Scores.Home),
					"period":       v.Status.Period,
					"timer":        v.Timer,
					"awayTeamName": v.Teams.Away.Name,
					"awayTeamLogo": v.Teams.Away.Logo,
					"awayScore":    strconv.Itoa(v.Scores.Away),
				}
				output = append(output, item)
			}
		}
	}
	utils.SetCache(cacheKey, output)

	return output, nil
}
