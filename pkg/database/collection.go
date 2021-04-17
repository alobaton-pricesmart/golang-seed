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
		db:    db.db.Model(model),
		model: model,
	}

	return c
}

// Pageable set up a Pageable request to SQL statement.
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

// Order st up a Sort request to SQL statement.
func (c *Collection) Order(sort Sort) *Collection {
	for _, s := range sort.sorters {
		c.db = c.db.Order(fmt.Sprintf("%v %v", s.field, s.direction))
	}
	return c
}

// Conditions set up conditions through a map or a Model to SQL statement.
func (c *Collection) Conditions(conditions interface{}) *Collection {
	mapc, ok := conditions.(map[string]interface{})
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

// Count perform a count.
func (c *Collection) Count(count *int64) error {
	result := c.db.Count(count)
	return result.Error
}

// Find perform a find all.
func (c *Collection) Find(instance interface{}) error {
	result := c.db.Find(instance)
	return result.Error
}

// Get perform a find one.
func (c *Collection) Get(instance Model) error {
	result := c.db.First(instance)
	return result.Error
}

// Update perform an update.
func (c *Collection) Update(instance Model) error {
	result := c.db.Updates(instance)
	return result.Error
}

// UpdateAll perform a batch update.
func (c *Collection) UpdateAll(instance interface{}) error {
	result := c.db.Updates(instance)
	return result.Error
}

// Create perform a create.
func (c *Collection) Create(instance Model) error {
	result := c.db.Create(instance)
	return result.Error
}

// CreateAll perform a batch create.
func (c *Collection) CreateAll(instance interface{}) error {
	result := c.db.Create(instance)
	return result.Error
}
