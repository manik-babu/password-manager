package user

import (
	"detrox/internal/auth"
	"detrox/internal/user/dto"
	"errors"
	"fmt"
)

type service struct {
	repo       *repository
	jwtService auth.JWTService
}

func NewService(repo *repository, jwtService auth.JWTService) *service {
	return &service{
		repo:       repo,
		jwtService: jwtService,
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
	token, jwtErr := s.jwtService.GenerateToken(user.ID, user.Email)
	if jwtErr != nil {
		fmt.Println(jwtErr)
		return nil, jwtErr
	}
	response := dto.Response{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
	return &response, nil
}
