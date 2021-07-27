package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/mail"
)

type Customer struct {
	gorm.Model
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Orders   []Order
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := Hash(c.Password)
	if err != nil {
		return err
	}
	c.Password = string(hashedPassword)
	return nil
}

func (c *Customer) Validate() error {
	if c.Password == "" {
		return errors.New("required password")
	}
	if c.Email == "" {
		return errors.New("required email")
	}
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return errors.New("invalid email")
	}
	return nil

}

func (c *Customer) CreateCustomer(db *gorm.DB) (*Customer, error) {

	var err error
	err = db.Model(&Customer{}).Create(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}

func (c *Customer) FindCustomerByID(db *gorm.DB, cid uint) (*Customer, error) {
	var err error
	err = db.Model(Customer{}).Where("id = ?", cid).First(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Customer{}, errors.New("customer not found")
	}
	return c, err
}
