package handler_test

import (
	"app/internal/app/apiserver"
	//"app/internal/app/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	databaseURL string = "root:112233@tcp(127.0.0.1:3306)/todolist?parseTime=true"
)

func Test_GetAllProject(t *testing.T) {
	//p := []*model.Project{}

	s := &apiserver.Server{}
	s.Initialize(databaseURL)

	testCases := []struct {
		name         string
		expectedCode int
	}{
		{
			name:         "valid",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/projects", nil)
			s.Router.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
