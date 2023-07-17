package request

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"widget-sports/utils"
)

func ServerRequest(host, key, url string, responseData interface{}) ([]byte, error) {

	timer := utils.MeasureTime()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("got: %w", err)
	}

	req.Header.Add("X-RapidAPI-Key", key)
	req.Header.Add("X-RapidAPI-Host", host)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("got: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("got: %w", err)
	}

	duration := timer()
	log.Printf("TimeRequest: %s\n", duration)

	err = json.Unmarshal(body, responseData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling bodyResponse: %w ", err)
	}

	return body, nil
}
