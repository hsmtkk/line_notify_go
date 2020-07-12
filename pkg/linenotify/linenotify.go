package linenotify

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Notifier interface {
	NotifyMessage(message string) error
	Status() error
}

type notifierImpl struct {
	client  *http.Client
	baseURL string
	token   string
}

const urlLINEAPI = "https://notify-api.line.me/api"

func New(token string) Notifier {
	return &notifierImpl{client: http.DefaultClient, baseURL: urlLINEAPI, token: token}
}

func (notifier *notifierImpl) NotifyMessage(message string) error {
	params := url.Values{}
	params.Add("message", message)
	req, err := http.NewRequest(http.MethodPost, notifier.baseURL+"/notify", strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+notifier.token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := notifier.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("got HTTP status code %d", resp.StatusCode)
}

func (notifier *notifierImpl) Status() error {
	req, err := http.NewRequest(http.MethodGet, notifier.baseURL+"/status", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+notifier.token)
	resp, err := notifier.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("got HTTP status code %d", resp.StatusCode)
}
