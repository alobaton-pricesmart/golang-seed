package database

import (
	"fmt"

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

func (c *Collection) Pageable(pageable Pageable) *Collection {
	if pageable.Page < 0 {
		pageable.Page = 0
	}

	if pageable.Size < 0 {
		pageable.Size = 10
	}

	c.db = c.db.Offset(pageable.Page).Limit(pageable.Size)
	return c
}

func (c *Collection) Order(sort Sort) *Collection {
	for _, s := range sort.sorters {
		c.db = c.db.Order(fmt.Sprintf("%v %v", s.field, s.direction))
	}
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

func (c *Collection) Find(instance interface{}) error {
	result := c.db.Find(instance)
	return result.Error
}

func (c *Collection) Get(instance Model) error {
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
