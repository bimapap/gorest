package service

import (
	"time"

	"github.com/bimapap/gorest/model"
	"github.com/bimapap/gorest/repository"
)

type CustomerService interface {
	GetCustomer(id int) (m *model.Customer, err error)
	GetAllCustomer(limit int, offset int) (model []model.Customer, err error)
	CreateCustomer(cust *model.Customer) error
	UpdateCustomer(id int, cust *model.Customer) error
	DeleteCustomer(id int) error
}

type customerService struct {
	customers      map[int]*model.Customer
	custRepository repository.CustomerRepository
}

func NewCustomerService(custRepository repository.CustomerRepository) CustomerService {
	return &customerService{
		custRepository: custRepository,
	}
}

func (c *customerService) getCustomerById(id int) (m *model.Customer, err error) {
	m, err = c.custRepository.FindOne(id)
	return
}

func (c *customerService) GetCustomer(id int) (m *model.Customer, err error) {
	return c.getCustomerById(id)
}
func (c *customerService) GetAllCustomer(limit int, offset int) (model []model.Customer, err error) {
	model, err = c.custRepository.FindAll(limit, offset)
	return
}

func (c *customerService) CreateCustomer(cust *model.Customer) error {
	cust.CreatedAt = time.Now()
	return c.custRepository.Create(cust)
}

func (c *customerService) UpdateCustomer(id int, cust *model.Customer) error {
	existingCustomer, err := c.getCustomerById(id)
	if err != nil {
		return err
	}

	var updateValue = model.Customer{
		Name:      cust.Name,
		Address:   cust.Address,
		UpdatedAt: time.Now(),
	}

	return c.custRepository.Update(existingCustomer, updateValue)
}

func (c *customerService) DeleteCustomer(id int) error {
	existingCustomer, err := c.getCustomerById(id)
	if err != nil {
		return err
	}

	var updateValue = model.Customer{
		DeletedAt: time.Now(),
	}

	return c.custRepository.Delete(existingCustomer, updateValue)
}
