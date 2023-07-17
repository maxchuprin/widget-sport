package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
	"widget-sports/configurations"
)

var (
	re = regexp.MustCompile(`^([\w\s]+) - ([\w\s]+)$`)
)

// сравнение содержимого
func Contains(options []string, searchValue string) bool {
	for _, option := range options {
		if strings.TrimSpace(option) == searchValue {
			return true
		}
	}
	return false
}

// выбор часового пояса
func CheckTimeZone(timeZone int) string {
	switch timeZone {
	case 0:
		return ""
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Europe/Istanbul"
	case 4:
		return ""
	case 5:
		return ""
	case 6:
		return "Asia/Almaty"
	case 7:
		return ""
	case 8:
		return ""
	case 9:
		return ""
	case 10:
		return ""
	case 11:
		return ""
	case 12:
		return ""
	case 13:
		return ""
	default:
		return "Asia/Almaty"
	}
}

// проверка на содержания цифр в имени игроков (API тенниса возвращают некорректные имена игроков)
func ContainsDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// получаю timeStamp, обрабатываю и возвращаю в формате: 2006-01-02T15:04:05-07:00
func DateFromTimeStamp(timeStamp int) (string, error) {
	ts := time.Unix(int64(timeStamp), 0)
	location, err := time.LoadLocation(CheckTimeZone(configurations.Config{}.TimeZone))
	if err != nil {
		return "", fmt.Errorf("error when processing a timeStamp in DateFromTimeStamp function: %w", err)
	}
	return ts.In(location).Format("2006-01-02T15:04:05-07:00"), nil
}

// регулярка для разделения строки на турнир и весовую категорию
func LineSplit(shortName string) (string, string) {
	var weightCategory, tournament string
	matches := re.FindStringSubmatch(shortName)
	if len(matches) > 2 {
		weightCategory = matches[1]
		tournament = matches[2]
	}
	return weightCategory, tournament
}

// время выполнения запроса
func MeasureTime() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		end := time.Now()
		return end.Sub(start)
	}
}

// сервис RAPID API возвращает сообщение о проблеме
func GetMessage(bodyResponse []byte) error {
	if strings.Contains(string(bodyResponse), "message") {
		return errors.New(string(bodyResponse))
	}
	return nil
}
