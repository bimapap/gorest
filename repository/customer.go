package repository

import (
	"github.com/bimapap/gorest/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(cust *model.Customer) error
	Update(cust *model.Customer, updateValue interface{}) error
	Delete(cust *model.Customer, updateValue interface{}) error
	FindOne(id int) (*model.Customer, error)
	FindAll(limit int, offset int) ([]model.Customer, error)
}

type customerRepository struct {
	dbConnection *gorm.DB
}

func NewCustomerRepository(dbConnection *gorm.DB) CustomerRepository {
	return &customerRepository{dbConnection: dbConnection}
}

func (c *customerRepository) Create(cust *model.Customer) error {
	return c.dbConnection.Create(cust).Error
}

func (c *customerRepository) FindOne(id int) (cust *model.Customer, err error) {
	cust = &model.Customer{}
	err = c.dbConnection.First(cust, id).Error
	return
}

func (c *customerRepository) FindAll(limit int, offset int) (cust []model.Customer, err error) {
	cust = []model.Customer{}
	err = c.dbConnection.Limit(limit).Offset(offset).Find(&cust).Error
	return
}

func (c *customerRepository) Update(cust *model.Customer, updateValue interface{}) error {
	return c.dbConnection.Model(cust).Updates(updateValue).Error
}
func (c *customerRepository) Delete(cust *model.Customer, updateValue interface{}) error {
	return c.dbConnection.Model(cust).Updates(updateValue).Error
}

// func (c *customerRepository) Delete(cust *model.Customer) error {
// 	return c.dbConnection.Delete(cust).Error
// }
