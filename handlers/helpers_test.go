package handlers

import (
	"net/url"
	"testing"
)

func TestValidateQuery(t *testing.T) {
	validValues := url.Values{
		"site":     []string{"cmore.se"},
		"video_id": []string{"123456"},
		"user_id":  []string{"123456"},
	}

	err := ValidateQuery(validValues)
	if err != nil {
		t.Fatal(err)
	}

	invalidValues := []url.Values{
		{
			"video_id": []string{"123456"},
			"user_id":  []string{"123456"},
		},
		{
			"site":    []string{"cmore.se"},
			"user_id": []string{"123456"},
		},
		{
			"site":     []string{"cmore.se"},
			"video_id": []string{"123456"},
		},
		{
			"site":     make([]string, 0),
			"video_id": []string{"123456"},
			"user_id":  []string{"123456"},
		},
		{
			"site":     []string{"cmore.se"},
			"video_id": make([]string, 0),
			"user_id":  []string{"123456"},
		},
		{
			"site":     []string{"cmore.se"},
			"video_id": []string{"123456"},
			"user_id":  make([]string, 0),
		},
	}
	for _, v := range invalidValues {
		err := ValidateQuery(v)
		if err == nil {
			t.Fatal("expecting error of missing query parameters")
		}
	}
}
