package tautulli

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStreamMetadata(t *testing.T) {
	streamMetadata := `{
		"group_count": null,
		"started": 1544312119,
		"original_title": "",
		"transcode_decision": "transcode",
		"parent_rating_key": "",
		"year": 2009,
		"paused_counter": 120,
		"duration": 11,
		"full_title": "Transformers: Revenge of the Fallen",
		"reference_id": null,
		"parent_title": "",
		"date": 1544312119,
		"percent_complete": 26,
		"ip_address": "24.4.181.217",
		"id": null,
		"session_key": 267,
		"group_ids": null,
		"user_id": 4070012,
		"thumb": "/library/metadata/18437/thumb/1543721875",
		"media_index": "",
		"player": "iPhone",
		"title": "",
		"friendly_name": "justinplexuser",
		"watched_status": 0,
		"rating_key": 18437,
		"platform": "iOS",
		"state": "paused",
		"stopped": 1544312250,
		"grandparent_title": "",
		"media_type": "movie",
		"parent_media_index": "",
		"grandparent_rating_key": "",
		"user": "justinplexuser"
	}`

	meta := StreamMetadata{}

	err := json.Unmarshal([]byte(streamMetadata), &meta)
	assert.Nil(t, err)
	assert.Equal(t, meta.Started, time.Unix(1544312119, 0))
	assert.Equal(t, meta.Transcoded, true)
	assert.Equal(t, meta.User, "justinplexuser")
	assert.Equal(t, meta.UserId, 4070012)
	assert.Equal(t, meta.MediaType, "movie")
	assert.Equal(t, meta.Player, "iPhone")
	assert.Equal(t, meta.Platform, "iOS")
}
