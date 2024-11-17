package memberships

import "music-app/internal/models/memberships"

func (r *repository) CreateUser(model memberships.User) error {
	return r.db.Create(&model).Error
}

func (r *repository) GetUser(email, username string, id uint) (*memberships.User, error) {
	var user memberships.User
	err := r.db.Where("email = ?", email).Or("username = ?", username).Or("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
