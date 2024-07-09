package user

import (
	"context"
	"os"
	"fmt"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	repo Respository
	timeout time.Duration
}

func NewService (repo Respository) *service {
	return &service{
		repo: repo,
		timeout: 5 * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel:=context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
	}

	r, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res:= &CreateUserRes{
		ID: r.ID,
		Username: r.Username,
		Email: r.Email,
	}

	return res, nil	
}

type MyJWTClaims struct {
	ID       string    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s* service) Login (ctx context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctxCancel, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	u, err := s.repo.GetUserByEmail(ctxCancel, req.Email)

	// 若是該Eamil尚未註冊，則呼叫CreateUser進行註冊，註冊完後繼續登入流程
	if err != nil {
		fmt.Println(err)
		return &LoginUserRes{}, fmt.Errorf("this email hasn't been registered")
	}

	err = util.CheckPasswordHash(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, fmt.Errorf("invalid password")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{
		accessToken: ss,
		Username: u.Username,
		Email: u.Email,
	}, nil
}
