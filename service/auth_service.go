package service

import (
	"9gag-api/contract"
	"9gag-api/repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
)

type AuthService struct {
	userRepository repository.UserRepositoryInterface
}

func NewAuthService(userRepository repository.UserRepositoryInterface) AuthService {
	return AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Signup(req *contract.SignupRequest) (*contract.SignupResponse, error) {
	if err := s.validateSignupRequest(req); err != nil {
		return nil, err
	}
	err := s.userRepository.Create(req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &contract.SignupResponse{}, nil
}

func (s *AuthService) Signin(req *contract.SigninRequest) (*contract.SigninResponse, error) {
	if err := s.validateSigninRequest(req); err != nil {
		return nil, err
	}
	user, err := s.userRepository.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	err = verifyPassword(user.Password, req.Password)
	if err != nil {
		return nil, err
	}
	token, err := generateToken(req.Username)
	if err != nil {
		return nil, err
	}
	return &contract.SigninResponse{
		Token: token,
	}, nil
}

func (s *AuthService) validateSignupRequest(req *contract.SignupRequest) error {
	if len(req.Username) == 0 {
		return errors.New("username cannot be empty")
	}
	if len(req.Email) == 0 {
		return errors.New("email cannot be empty")
	}
	if len(req.Password) == 0 {
		return errors.New("password cannot be empty")
	}
	user, err := s.userRepository.GetByUsernameOrEmail(req.Username, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	if user.Username == req.Username {
		return errors.New("username already exist")
	}
	if user.Email == req.Email {
		return errors.New("email already exist")
	}
	return nil
}

func (s *AuthService) validateSigninRequest(req *contract.SigninRequest) error {
	if len(req.Username) == 0 {
		return errors.New("username cannot be empty")
	}
	if len(req.Password) == 0 {
		return errors.New("password cannot be empty")
	}
	return nil
}

func verifyPassword(userPassword string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return errors.New("incorrect password")
	}
	return nil
}

func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

