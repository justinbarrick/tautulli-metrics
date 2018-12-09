package metrics

import (
	influxdb "github.com/influxdata/influxdb/client/v2"
	"github.com/justinbarrick/tautulli-metrics/pkg/tautulli"
	"strconv"
	"time"
)

type Metrics struct {
	influx   influxdb.Client
	database string
}

func NewMetrics(address, database, user, pass string) (*Metrics, error) {
	m := Metrics{}

	influx, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr:     address,
		Username: user,
		Password: pass,
	})
	if err != nil {
		return nil, err
	}

	m.influx = influx
	m.database = database
	return &m, nil
}

func (m Metrics) MostRecentHistoryTimestamp() (time.Time, error) {
	response, err := m.influx.Query(influxdb.Query{
		Command:  `SELECT last("title"), time FROM history`,
		Database: m.database,
	})
	if err != nil {
		return time.Time{}, nil
	}
	if response.Error() != nil {
		return time.Time{}, nil
	}

	lastTimestampStr := response.Results[0].Series[0].Values[0][0].(string)
	return time.Parse(time.RFC3339, lastTimestampStr)
}

func (m Metrics) InsertHistory(streams []tautulli.StreamMetadata) error {
	bp, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
		Database:  m.database,
		Precision: "ms",
	})
	if err != nil {
		return err
	}

	for _, stream := range streams {
		pt, err := influxdb.NewPoint("history", map[string]string{
			"user":       stream.User,
			"media-type": stream.MediaType,
			"transcoded": strconv.FormatBool(stream.Transcoded),
		}, map[string]interface{}{
			"title":    stream.Title,
			"userId":   stream.UserId,
			"player":   stream.Player,
			"platform": stream.Platform,
		}, stream.Started)
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
	}

	if err := m.influx.Write(bp); err != nil {
		return err
	}

	return nil
}
