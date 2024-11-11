package memberships

import "time"

type (
	SignUpRequest struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `json:"refreshToken" binding:"required"`
	}
)

type (
	LoginResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}

	RefreshTokenResponse struct {
		AccessToken string `json:"accessToken"`
	}
)

type (
	UserModel struct {
		ID        int64     `db:"id"`
		Username  string    `db:"username"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedBy string    `db:"updated_by"`
	}

	RefreshTokenModel struct {
		ID           int64     `db:"id"`
		UserID       int64     `db:"user_id"`
		RefreshToken string    `db:"refresh_token"`
		ExpiredAt    time.Time `db:"expired_at"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
		CreatedBy    string    `db:"created_by"`
		UpdatedBy    string    `db:"updated_by"`
	}
)
