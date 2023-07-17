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

type dataTennis struct {
	Data []structures.TennisMatch `json:"data"`
}

func GetTennisMatchesByDate(cfg configurations.Config) ([]map[string]interface{}, error) {
	cacheKey := "tennisByDate"
	matches, found := utils.GetFromCache(cacheKey)
	if found {
		return matches, nil
	}

	url := fmt.Sprintf("https://%s/sports/2/events/date/%s", cfg.TennisHost, time.Now().AddDate(0, 0, cfg.PlusDay).Format("2006-01-02"))

	var d dataTennis

	bodyResponse, err := request.ServerRequest(cfg.TennisHost, cfg.RapidAPIKey, url, &d)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	err = utils.GetMessage(bodyResponse)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	var output []map[string]interface{}
	for _, v := range d.Data {
		if utils.Contains(strings.Split(cfg.TennisLeague, ","), v.League.Name) && !utils.ContainsDigit(v.Home.Name) && !utils.ContainsDigit(v.Away.Name) {
			t, _ := time.Parse("2006-01-02 15:04:05", v.Date)
			t = t.In(time.FixedZone("UTC+6", cfg.TimeZone*60*60))
			matchDate := t.Format("2006-01-02T15:04:05-07:00")
			item := map[string]interface{}{
				"matchDateTime": matchDate,
				"leagueName":    v.League.Name,
				"leagueLogo":    v.League.Logo,
				"homeTeamName":  v.Home.Name,
				"homeTeamLogo":  v.Home.Logo,
				"awayTeamName":  v.Away.Name,
				"awayTeamLogo":  v.Away.Logo,
				"stage":         v.Stage.Name,
			}
			output = append(output, item)
		}
	}
	utils.SetCache(cacheKey, output)

	return output, nil
}

func GetTennisMatchesByLiveScore(cfg configurations.Config) ([]map[string]interface{}, error) {
	cacheKey := "tennisLiveScore"
	matches, found := utils.GetFromCache(cacheKey)
	if found {
		return matches, nil
	}

	url := fmt.Sprintf("https://%s/sports/2/events/live", cfg.TennisHost)

	var d dataTennis
	bodyResponse, err := request.ServerRequest(cfg.TennisHost, cfg.RapidAPIKey, url, &d)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	err = utils.GetMessage(bodyResponse)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	var output []map[string]interface{}
	for _, v := range d.Data {
		if utils.Contains(strings.Split(cfg.TennisLeague, ","), v.League.Name) && !utils.ContainsDigit(v.Home.Name) && !utils.ContainsDigit(v.Away.Name) {
			item := map[string]interface{}{
				"leagueName":   v.League.Name,
				"leagueLogo":   v.League.Logo,
				"homeTeamName": v.Home.Name,
				"homeTeamLogo": v.Home.Logo,
				"homeScore":    v.HomeScore,
				"awayTeamName": v.Away.Name,
				"awayTeamLogo": v.Away.Logo,
				"awayScore":    v.AwayScore,
				"stage":        v.Stage.Name,
			}
			output = append(output, item)
		}
	}
	utils.SetCache(cacheKey, output)

	return output, nil
}
