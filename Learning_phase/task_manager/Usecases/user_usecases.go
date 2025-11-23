package usecases

import (
	"errors"
	"task_manager/Domain"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
)

type UserUsecase interface {
	Register(username, password string) (domain.User, error)
	Login(username, password string) (domain.User, string, error)
	PromoteUser(username string) error
}

type userUsecase struct {
	userRepo        repositories.UserRepository
	passwordService *infrastructure.PasswordService
	jwtService      *infrastructure.JWTService
}

func NewUserUsecase(userRepo repositories.UserRepository, passwordService *infrastructure.PasswordService, jwtService *infrastructure.JWTService) UserUsecase {
	return &userUsecase{
		userRepo:        userRepo,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (u *userUsecase) Register(username, password string) (domain.User, error) {
	existing, _ := u.userRepo.GetByUsername(username)
	if existing.Username != "" {
		return domain.User{}, errors.New("username already exists")
	}

	count, err := u.userRepo.CountUsers()
	if err != nil {
		return domain.User{}, err
	}

	role := "user"
	if count == 0 {
		role = "admin"
	}

	hashedPassword, err := u.passwordService.HashPassword(password)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	createdUser, err := u.userRepo.Create(user)
	if err != nil {
		return domain.User{}, err
	}

	createdUser.Password = ""
	return createdUser, nil
}

func (u *userUsecase) Login(username, password string) (domain.User, string, error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return domain.User{}, "", errors.New("invalid credentials")
	}

	err = u.passwordService.ComparePassword(user.Password, password)
	if err != nil {
		return domain.User{}, "", errors.New("invalid credentials")
	}

	token, err := u.jwtService.GenerateToken(user.ID.Hex(), user.Username, user.Role)
	if err != nil {
		return domain.User{}, "", err
	}

	user.Password = ""
	return user, token, nil
}

func (u *userUsecase) PromoteUser(username string) error {
	return u.userRepo.PromoteToAdmin(username)
}
