package memberships

import (
	"music-app/internal/configs"
	"music-app/internal/models/memberships"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
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

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		req memberships.LoginRequest
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
				req: memberships.LoginRequest{
					Email:    "test@test.com",
					Password: "111111",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("111111"), bcrypt.DefaultCost)
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@test.com",
					Password: string(hashedPassword),
					Username: "raffyjo",
				}, nil)
			},
		},
		{
			name: "failed when get user",
			args: args{
				req: memberships.LoginRequest{
					Email:    "test@test.com",
					Password: "111111",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when password is wrong",
			args: args{
				req: memberships.LoginRequest{
					Email:    "test@test.com",
					Password: "111111",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@test.com",
					Password: "wrong password",
					Username: "raffyjo",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJWT: "SECRET",
					},
				},
				repository: mockRepo,
			}
			got, err := s.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
