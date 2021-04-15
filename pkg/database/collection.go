package database

import (
	"gorm.io/gorm"
)

// Collection relies on gorm, if don't want to use gorm you must provide your own implementation of Collection's methods.
type Collection struct {
	db    *gorm.DB
	model Model
}

func newCollection(db *Database, model Model) *Collection {
	c := &Collection{
		db:    db.db,
		model: model,
	}

	return c
}

func (c *Collection) Pageable(page, size int) *Collection {
	if page < 0 {
		page = 0
	}

	if size < 0 {
		size = 10
	}

	c.db = c.db.Offset(page).Limit(size)
	return c
}

func (c *Collection) Conditions(conditions interface{}) *Collection {
	mapc, ok := conditions.(*map[string]string)
	if ok {
		c.db = c.db.Where(mapc)
		return c
	}

	modc, ok := conditions.(Model)
	if ok {
		c.db = c.db.Where(modc)
		return c
	}

	return c
}

func (c *Collection) FindAll(instance interface{}) error {
	result := c.db.Find(instance)
	return result.Error
}

func (c *Collection) GetByID(instance Model) error {
	result := c.db.First(instance)
	return result.Error
}

func (c *Collection) Update(instance Model) error {
	result := c.db.Updates(instance)
	return result.Error
}

func (c *Collection) UpdateAll(instance interface{}) error {
	result := c.db.Updates(instance)
	return result.Error
}

func (c *Collection) Create(instance Model) error {
	result := c.db.Create(instance)
	return result.Error
}

func (c *Collection) CreateAll(instance interface{}) error {
	result := c.db.Create(instance)
	return result.Error
}
