package model_test

import (
	"testing"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestRound_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		r       func() *model.Round
		isValid bool
	}{
		{
			name: "valid",
			r: func() *model.Round {
				return model.TestRound(t)
			},
			isValid: true,
		},
		{
			name: "invalid",
			r: func() *model.Round {
				r := &model.Round{
					UserID:     0,
					Difficulty: 0,
					GameNumber: "GFAR",
					GameTime:   "",
					Inpt:       "GFDS",
					Outpt:      "9876",
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
