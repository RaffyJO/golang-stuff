package memberships

import (
	"music-app/internal/models/memberships"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		req memberships.SignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				req: memberships.SignUpRequest{
					Email:    "test@test.com",
					Username: "test",
					Password: "test",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, args.req.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
		{
			name: "failed when get user",
			args: args{
				req: memberships.SignUpRequest{
					Email:    "test@test.com",
					Username: "test",
					Password: "test",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, args.req.Username, uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when create user",
			args: args{
				req: memberships.SignUpRequest{
					Email:    "test@test.com",
					Username: "test",
					Password: "test",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, args.req.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				repository: mockRepo,
			}
			if err := s.SignUp(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("service.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
