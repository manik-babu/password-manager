package user

import "detrox/internal/user/dto"

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(req dto.CreateUserRequest) (*dto.Response, error) {
	user := User{
		Name:  req.Name,
		Email: req.Email,
	}
	// hashing user password with bcrypt
	user.hashPassword(req.Password)

	// calling repository layer here
	err := s.repo.CreateUser(&user)

	if err != nil {
		return nil, err
	}
	response := dto.Response{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return &response, nil
}
