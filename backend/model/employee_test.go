package model_test

import (
	"testing"
	"time"

	"github.com/matheusmosca/nutcache/model"
	"github.com/stretchr/testify/assert"
)

func TestNewEmployee(t *testing.T) {
	e, err := model.NewEmployee("George", "George@g.com", "male", "Backend", "11122233312", time.Now(), time.Now())
	assert.Equal(t, e.Name, "George")
	assert.Equal(t, e.Email, "George@g.com")
	assert.Equal(t, e.Gender, "male")
	assert.Equal(t, e.Team, "Backend")
	assert.Equal(t, e.CPF, "11122233312")

	assert.NotNil(t, e.ID)
	assert.NotNil(t, e.BirthDate)
	assert.NotNil(t, e.StartDate)
	assert.NotNil(t, e.CreatedAt)
	assert.NotNil(t, e.UpdatedAt)

	assert.Nil(t, err)
}

func TestValidate(t *testing.T) {
	type test struct {
		name      string
		email     string
		gender    string
		CPF       string
		team      string
		birthDate time.Time
		startDate time.Time
		want      error
	}

	tests := []test{
		{
			name:      "",
			email:     "em@em.com",
			gender:    "male",
			CPF:       "22233322214",
			birthDate: time.Now(),
			startDate: time.Now(),
			team:      "Frontend",
			want:      model.ErrInvalidName,
		},
		{
			name:      "emily",
			email:     "",
			gender:    "male",
			CPF:       "22233322214",
			birthDate: time.Now(),
			startDate: time.Now(),
			team:      "Frontend",
			want:      model.ErrInvalidEmail,
		},
		{
			name:      "emily",
			email:     "em@em.com",
			gender:    "",
			CPF:       "22233322214",
			birthDate: time.Now(),
			startDate: time.Now(),
			team:      "Frontend",
			want:      model.ErrInvalidGender,
		},
		{
			name:      "emily",
			email:     "em@em.com",
			gender:    "male",
			CPF:       "123213",
			birthDate: time.Now(),
			startDate: time.Now(),
			team:      "Frontend",
			want:      model.ErrInvalidCPF,
		},
		{
			name:      "emily",
			email:     "em@em.com",
			gender:    "male",
			CPF:       "22233322214",
			birthDate: time.Now(),
			startDate: time.Now(),
			team:      "Frontend",
			want:      nil,
		},
	}
	for _, tt := range tests {
		_, err := model.NewEmployee(tt.name, tt.email, tt.gender, tt.team, tt.CPF, tt.birthDate, tt.startDate)
		assert.Equal(t, err, tt.want)
	}
}
