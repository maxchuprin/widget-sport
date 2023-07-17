package services

import (
	"fmt"
	"strconv"
	"strings"
	"widget-sports/configurations"
	"widget-sports/request"
	"widget-sports/structures"
	"widget-sports/utils"
)

type dataMMA struct {
	Data []structures.Data `json:"DATA"`
}

func GetMMAMatchesByDate(cfg configurations.Config) ([]map[string]interface{}, error) {
	cacheKey := "mmaByDate"
	matches, found := utils.GetFromCache(cacheKey)
	if found {
		return matches, nil
	}

	url := fmt.Sprintf("https://%s/v1/events/list?timezone=%s&indent_days=%s&locale=en_INT&&sport_id=28", cfg.MMAHost, strconv.Itoa(cfg.TimeZone), strconv.Itoa(cfg.PlusDay))

	var d dataMMA

	bodyResponse, err := request.ServerRequest(cfg.MMAHost, cfg.RapidAPIKey, url, &d)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	err = utils.GetMessage(bodyResponse)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	jsonResultByMatches, err := jsonProcessing(&d)
	if err != nil {
		return nil, fmt.Errorf("got: %w", err)
	}
	for _, j := range jsonResultByMatches {
		delete(j, "winner")
	}
	utils.SetCache(cacheKey, jsonResultByMatches)
	return jsonResultByMatches, nil
}

func GetMMAMatchesResults(cfg configurations.Config) ([]map[string]interface{}, error) {
	cacheKey := "mmaResults"
	matches, found := utils.GetFromCache(cacheKey)
	if found {
		return matches, nil
	}

	url := fmt.Sprintf("https://%s/v1/events/list?timezone=%s&indent_days=-1&locale=en_INT&sport_id=28", cfg.MMAHost, strconv.Itoa(cfg.TimeZone))

	var d dataMMA
	bodyResponse, err := request.ServerRequest(cfg.MMAHost, cfg.RapidAPIKey, url, &d)

	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	err = utils.GetMessage(bodyResponse)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}

	jsonResultByMatches, err := jsonProcessing(&d)
	if err != nil {
		return nil, fmt.Errorf("wrap error: %w", err)
	}
	utils.SetCache(cacheKey, jsonResultByMatches)
	return jsonResultByMatches, nil
}

func jsonProcessing(d *dataMMA) ([]map[string]interface{}, error) {
	var output []map[string]interface{}
	var homeImage, awayImage string
	for _, v := range d.Data {
		if strings.Contains(v.ShortName, "UFC Men") && v.CategoryName == "World" {
			for _, e := range v.Events {
				if e.HomeImages != nil {
					homeImage = e.HomeImages[0]
				}
				if e.AwayImages != nil {
					awayImage = e.AwayImages[0]
				}
				matchDate, err := utils.DateFromTimeStamp(e.TimeStamp)
				if err != nil {
					return nil, fmt.Errorf("got: %w", err)
				}
				weightCategory, tournament := utils.LineSplit(v.ShortName)

				item := map[string]interface{}{
					"matchDateTime":  matchDate,
					"tournament":     tournament,
					"weightCategory": weightCategory,
					"winner":         strconv.Itoa(e.Winner),
					"homeTeamName":   e.HomeName,
					"homeTeamLogo":   homeImage,
					"awayTeamName":   e.AwayName,
					"awayTeamLogo":   awayImage,
				}
				output = append(output, item)
			}
		}
	}
	return output, nil
}
