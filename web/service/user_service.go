package service

import (
	"context"
    "myproject/web/model"
    "myproject/web/repository"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"fmt"
	"myproject/web/utils"
)

type UserService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, id uint) (*model.User, error) {
	return s.repo.FindById(ctx, id)
}

func (s *UserService) Register(ctx context.Context, user model.User) (*model.User, error) {

	// 加密密码
	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return nil, err
	}

	user.Password = string(hashedPassword)

	if _,err = s.repo.Create(ctx, &user); err != nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return nil, err
	}
	// ctx.JSON(http.StatusCreated, gin.H{"message" : "User registered successfully"})
	return &user, nil
}

func (s *UserService) Login(ctx context.Context, user model.User) (string, error) {
	// var storeUser model.User
	storeUser,err := s.repo.FindByName(ctx, user.Username); 
	if err != nil || storeUser == nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return "", errors.New("not found user")
	}

	fmt.Printf("oldpwd:%v, newpwd: %v", storeUser, user.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(storeUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	
	// 生成 jwt token
	tokenString, err := utils.GenerateToken(*storeUser)
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"id":	storeUser.ID,
	// 	"username": storeUser.Username,
	// 	"exp": time.Now().Add(time.Hour * 24).Unix(),
	// })

	// tokenString, err  := token.SignedString([]byte("xxoxx"))
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return tokenString, nil
}