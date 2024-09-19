package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetPrice() float64
	GetStatus() string
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID string `valid:"uuidv4"`
	Name string `valid:"required"`
	Price float64 `valid:"float, optional"`
	Status string `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("The status must be enabled or disabled");
	}

	if p.Price < 0 {
		return false, errors.New("The price must be a positive number")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("The price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
    	return nil
	}

	return errors.New("The price must be zero to disable the product")
}

func (p *Product) GetID() string {
    return p.ID
}

func (p *Product) GetName() string {
    return p.Name
}

func (p *Product) GetPrice() float64 {
    return p.Price
}

func (p *Product) GetStatus() string {
    return p.Status
}