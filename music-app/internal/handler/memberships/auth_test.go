package memberships

import (
	"bytes"
	"encoding/json"
	"music-app/internal/models/memberships"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_SignUp(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSvc := NewMockservice(mockCtrl)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@test.com",
					Username: "test",
					Password: "test",
				}).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "error",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@test.com",
					Username: "test",
					Password: "test",
				}).Return(assert.AnError)
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()

			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoutes()
			w := httptest.NewRecorder()

			endpoint := `/memberships/signup`
			model := memberships.SignUpRequest{
				Email:    "test@test.com",
				Username: "test",
				Password: "test",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
