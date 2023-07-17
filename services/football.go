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

type respFootball struct {
	Response []structures.FootballMatch `json:"response"`
}

func GetFootballMatchesByDate(cfg configurations.Config) ([]map[string]interface{}, error) {
	cacheKey := "footballByDate"
	matches, found := utils.GetFromCache(cacheKey)
	if found {
		return matches, nil
	}

	url := fmt.Sprintf("https://%s/v3/fixtures?date=%s&timezone=%s", cfg.FootballHost, time.Now().AddDate(0, 0, cfg.PlusDay).Format("2006-01-02"), utils.CheckTimeZone(cfg.TimeZone))

	var r respFootball

	bodyResponse, err := request.ServerRequest(cfg.FootballHost, cfg.RapidAPIKey, url, &r)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	err = utils.GetMessage(bodyResponse)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	var output []map[string]interface{}
	for _, v := range r.Response {
		item := map[string]interface{}{
			"matchDateTime": v.Fixture.Date,
			"leagueName":    v.League.Name,
			"leagueLogo":    v.League.Logo,
			"homeTeamName":  v.Teams.Home.Name,
			"homeTeamLogo":  v.Teams.Home.Logo,
			"awayTeamName":  v.Teams.Away.Name,
			"awayTeamLogo":  v.Teams.Away.Logo,
		}
		if (v.League.Name == "UEFA Champions League" && !strings.Contains(v.League.Round, "Qualifying")) || v.League.Name == "UEFA Europa League" {
			output = append(output, item)
		} else if (utils.Contains(strings.Split(cfg.FootballLeague, ","), v.League.Name)) && (utils.Contains(strings.Split(cfg.FootballTeam, ","), v.Teams.Home.Name) || utils.Contains(strings.Split(cfg.FootballTeam, ","), v.Teams.Away.Name)) {
			output = append(output, item)
		}
	}
	utils.SetCache(cacheKey, output)

	return output, nil
}

func GetFootballMatchesByLiveScore(cfg configurations.Config) ([]map[string]interface{}, error) {
	cacheKey := "footballLiveScore"
	matches, found := utils.GetFromCache(cacheKey)
	if found {
		return matches, nil
	}

	url := fmt.Sprintf("https://%s/v3/fixtures?live=all&timezone=%s", cfg.FootballHost, utils.CheckTimeZone(cfg.TimeZone))

	var r respFootball

	bodyResponse, err := request.ServerRequest(cfg.FootballHost, cfg.RapidAPIKey, url, &r)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	err = utils.GetMessage(bodyResponse)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	var output []map[string]interface{}
	var homeGoal, awayGoal map[string]string
	var homeGoals, awayGoals []map[string]string
	for _, v := range r.Response {
		for _, j := range v.Events {
			if j.Type == "Goal" {
				if strings.Contains(v.Teams.Home.Name, j.Team.Name) {
					homeGoal = map[string]string{
						"minuteGoal": strconv.Itoa(j.Time.Minute),
						"playerName": j.Player.Name,
					}
					homeGoals = append(homeGoals, homeGoal)
				} else if strings.Contains(v.Teams.Away.Name, j.Team.Name) {
					awayGoal = map[string]string{
						"minuteGoal": strconv.Itoa(j.Time.Minute),
						"playerName": j.Player.Name,
					}
					awayGoals = append(awayGoals, awayGoal)
				}
			}
		}
		item := map[string]interface{}{
			"leagueName":   v.League.Name,
			"leagueLogo":   v.League.Logo,
			"homeTeamName": v.Teams.Home.Name,
			"homeTeamLogo": v.Teams.Home.Logo,
			"homeScore":    strconv.Itoa(v.Score.Home),
			"homeGoals":    homeGoals,
			"minutesGame":  strconv.Itoa(v.Fixture.Status.MinuteGame),
			"awayTeamName": v.Teams.Away.Name,
			"awayTeamLogo": v.Teams.Away.Logo,
			"awayScore":    strconv.Itoa(v.Score.Away),
			"awayGoals":    awayGoals,
		}
		homeGoals = nil
		awayGoals = nil

		if (v.League.Name == "UEFA Champions League" && !strings.Contains(v.League.Round, "Qualifying")) || v.League.Name == "UEFA Europa League" {
			output = append(output, item)
		} else if (utils.Contains(strings.Split(cfg.FootballLeague, ","), v.League.Name)) &&
			(utils.Contains(strings.Split(cfg.FootballTeam, ","), v.Teams.Home.Name) || utils.Contains(strings.Split(cfg.FootballTeam, ","), v.Teams.Away.Name)) {
			output = append(output, item)
		}
	}
	utils.SetCache(cacheKey, output)

	return output, nil
}
