package pgsql

import (
	"context"
	"go-community/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (err error)
	GetByAccountNumber(ctx context.Context, accountNumber string) (user models.User, err error)
	GetByEmail(ctx context.Context, email string) (user models.User, err error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (user models.User, err error)
}

type userRepository struct {
	db  *gorm.DB
	trx TransactionRepository
}

func NewUserRepository(db *gorm.DB, trx TransactionRepository) UserRepository {
	return &userRepository{db: db, trx: trx}
}

func (ur *userRepository) Create(ctx context.Context, user *models.User) (err error) {
	defer func() {
		LogRepository(ctx, err)
	}()

	return ur.trx.Transaction(func(dtx *gorm.DB) error {
		return ur.db.Create(&user).Error
	})
}

func (ur *userRepository) GetByAccountNumber(ctx context.Context, accountNumber string) (user models.User, err error) {
	defer func() {
		LogRepository(ctx, err)
	}()

	var u models.User
	err = ur.db.Where("account_number = ?", accountNumber).First(&u).Error

	return u, err
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (user models.User, err error) {
	defer func() {
		LogRepository(ctx, err)
	}()

	var u models.User
	err = ur.db.Where("email = ?", email).First(&u).Error

	return u, err
}

func (ur *userRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (user models.User, err error) {
	defer func() {
		LogRepository(ctx, err)
	}()

	var u models.User
	err = ur.db.Where("phone_number = ?", phoneNumber).First(&u).Error

	return u, err
}
