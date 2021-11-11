package server

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRoles(t *testing.T) {
	repo := stubRepository{}
	s := apiServer{
		app:  fiber.New(),
		repo: &repo,
	}
	s.addRoleRoutes()

	tests := []struct {
		method   string
		endPoint string
		body     string
		code     int
		want     string
	}{
		{
			method:   "GET",
			endPoint: "/role",
			body:     "",
			code:     200,
			want:     defaultRolesRespStr,
		},
		{
			method:   "GET",
			endPoint: "/role/engineer",
			body:     "",
			code:     200,
			want:     defaultRoleRespStr,
		},
		{
			method:   "GET",
			endPoint: "/role/none",
			body:     "",
			code:     404,
			want:     notFoundStr,
		},
		{
			method:   "POST",
			endPoint: "/role",
			body:     `{"name": "myrole"}`,
			code:     201,
			want:     "",
		},
		{
			method:   "POST",
			endPoint: "/role",
			body:     "",
			code:     400,
			want:     "",
		},
		{
			method:   "DELETE",
			endPoint: "/role/1",
			body:     "",
			code:     300,
			want:     "",
		},
	}

	for _, tc := range tests {
		err := repo.Reset()
		assert.Nil(t, err)

		req := httptest.NewRequest(tc.method, tc.endPoint, bytes.NewBuffer([]byte(tc.body)))

		resp, err := s.app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, tc.code, resp.StatusCode)

		if tc.want != "" {
			got, err := getResponseText(t, resp)
			assert.Nil(t, err)

			assert.Equal(t, tc.want, got)
		}
	}
}

func getResponseText(t *testing.T, resp *http.Response) (string, error) {
	t.Helper()

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)

	return string(b), err
}
