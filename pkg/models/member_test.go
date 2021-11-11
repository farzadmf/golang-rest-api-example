package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	cases := []struct {
		member    Member
		wantValid bool
		message   string
		wantErr   error
	}{
		{
			member: Member{
				Type: Contractor,
			},
			wantValid: false,
			message:   "contractor without duration is not valid",
			wantErr:   ErrDurationRequired,
		},
		{
			member: Member{
				Type: Employee,
			},
			wantValid: false,
			message:   "employee without a role is not valid",
			wantErr:   ErrRoleRequired,
		},
		{
			member: Member{
				Type:     Contractor,
				Duration: 5,
			},
			wantValid: true,
			message:   "valid contractor",
			wantErr:   nil,
		},
		{
			member: Member{
				Type: Employee,
				Role: &Role{},
			},
			wantValid: true,
			message:   "valid employee",
			wantErr:   nil,
		},
	}

	for _, test := range cases {
		got, err := test.member.IsValid()

		assert.Equal(t, test.wantErr, err)
		assert.Equal(t, test.wantValid, got, test.message)
	}
}
