package model_test

import (
	"testing"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestRecord_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		r       func() *model.Record
		isValid bool
	}{
		{
			name: "valid",
			r: func() *model.Record {
				return model.TestRecord(t)
			},
			isValid: true,
		},
		{
			name: "invalid",
			r: func() *model.Record {
				r := &model.Record{
					UserID:     0,
					Difficulty: 0,
					RoundCount: 0,
					GameTime:   0,
				}
				return r
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.r().Validate())
			} else {
				assert.Error(t, tc.r().Validate())
			}
		})
	}
}
