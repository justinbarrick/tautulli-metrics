package tautulli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
)

type TautulliAPI struct {
	apiKey string
	apiUrl string
	client *http.Client
}

type TautulliAPIHeader struct {
	Response TautulliAPIResponse `json:"response"`
}

type TautulliAPIResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Result  string      `json:"result"`
}

type TautulliAPIHistory struct {
	Data            []StreamMetadata `json"data"`
	Draw            int              `json:"draw"`
	TotalDuration   string           `json:"total_duration"`
	FilterDuration  string           `json:"filter_duration"`
	RecordsFiltered int              `json:"recordsFiltered"`
	RecordsTotal    int              `json:"recordsTotal"`
}

type TautulliAPIActivity struct {
	Sessions                []StreamMetadata `json:"sessions"`
	LanBandwidth            int              `json:"lan_bandwidth"`
	StreamCount             string           `json:"stream_count"`
	StreamCountDirectPlay   int              `json:"stream_count_direct_play"`
	StreamCountDirectStream int              `json:"stream_count_direct_stream"`
	StreamCountTranscode    int              `json:"stream_count_transcode"`
	TotalBandwidth          int              `json:"total_bandwidth"`
	WanBandwidth            int              `json:"wan_bandwidth"`
}

func NewTautulliAPI(apiUrl, apiKey string) *TautulliAPI {
	return &TautulliAPI{
		client: &http.Client{},
		apiUrl: apiUrl,
		apiKey: apiKey,
	}
}

func (t TautulliAPI) Request(command string, options map[string]string, data interface{}) error {
	apiUrl, err := url.Parse(t.apiUrl)
	if err != nil {
		return err
	}

	apiUrl.Path = filepath.Join(apiUrl.Path, "/api/v2")

	values := url.Values{}
	values.Set("cmd", command)
	values.Set("apikey", t.apiKey)
	for key, value := range options {
		values.Set(key, value)
	}

	apiUrl.RawQuery = values.Encode()

	resp, err := t.client.Get(apiUrl.String())
	if err != nil {
		return err
	}

	response := TautulliAPIHeader{
		Response: TautulliAPIResponse{
			Data: data,
		},
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	return nil
}

func (t TautulliAPI) GetHistory(after time.Time) ([]StreamMetadata, error) {
	streams := []StreamMetadata{}

	index := 0
	more := true
	length := 500

	for more {
		history := TautulliAPIHistory{}
		err := t.Request("get_history", map[string]string{
			"length":       fmt.Sprintf("%d", length),
			"order_column": "date",
			"order_dir":    "desc",
			"start":        fmt.Sprintf("%d", index),
		}, &history)
		if err != nil {
			return streams, err
		}

		for _, stream := range history.Data {
			if stream.Started.Sub(after).Seconds() <= 0 {
				more = false
				break
			}

			streams = append(streams, stream)
		}

		if len(history.Data) < length {
			break
		}

		index += length
	}

	return streams, nil
}

func (t TautulliAPI) GetActivity() error {
	activity := TautulliAPIActivity{}

	err := t.Request("get_activity", map[string]string{}, &activity)
	if err != nil {
		return err
	}

	return nil
}
