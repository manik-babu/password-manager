package user

import (
	"detrox/internal/user/dto"
	"errors"
)

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
func (s *service) LoginUser(req dto.LoginUserRequest) (*dto.Response, error) {
	user, err := s.repo.GetUserByEmail(req.Email)

	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Email or password is incorrect")
	}

	checkPasswordError := user.checkPassword(req.Password)
	if checkPasswordError != nil {
		return nil, errors.New("Email or password is incorrect")
	}
	response := dto.Response{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return &response, nil

}
