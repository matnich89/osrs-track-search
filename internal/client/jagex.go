package client

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type HighscoresType string

const (
	Normal   HighscoresType = "hiscore_oldschool"
	Ironman  HighscoresType = "hiscore_oldschool_ironman"
	Hardcore HighscoresType = "hiscore_oldschool_hardcore_ironman"
	Ultimate HighscoresType = "hiscore_oldschool_ultimate"
)

//go:generate mockgen -destination=mocks/mock_client.go -package=mocks . Client
type Client interface {
	GetHighScores(character string, highscoresType HighscoresType) (resp *http.Response, err error)
}

var ErrNotFound = errors.New("not found")

type JagexClient struct {
	httpClient *http.Client
	baseUrl    string
}

func NewJagexClient(host string) (Client, error) {
	_, err := url.Parse(host)

	if err != nil {
		return nil, errors.New("invalid url")
	}

	baseUrl := host + "/m=%s/index_lite.ws?player=%s"

	return &JagexClient{
		httpClient: &http.Client{},
		baseUrl:    baseUrl,
	}, nil

}

func (c *JagexClient) GetHighScores(character string, highscoresType HighscoresType) (resp *http.Response, err error) {
	highscoresUrl := fmt.Sprintf(c.baseUrl, highscoresType, character)

	log.Println(highscoresUrl)

	req, err := http.NewRequest(http.MethodGet, highscoresUrl, nil)
	if err != nil {
		return nil, errors.New("could not create request")
	}

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return nil, errors.New("could not make request")
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, errors.New("bad response from jagex")
	}

	return resp, nil
}
