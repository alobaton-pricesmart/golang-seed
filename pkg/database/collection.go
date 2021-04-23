package database

import (
	"errors"
	"fmt"
	"golang-seed/pkg/pagination"
	"golang-seed/pkg/sorting"

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
func (c *Collection) Pageable(pageable pagination.Pageable) *Collection {
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
func (c *Collection) Order(sort sorting.Sort) *Collection {
	for _, s := range sort.Sorters {
		c.db = c.db.Order(fmt.Sprintf("%v %v", s.Field, s.Direction))
	}
	return c
}

// Conditions set up conditions through a map or a Model to SQL statement.
func (c *Collection) Conditions(conditions interface{}) *Collection {
	mapc, ok := conditions.(map[string]interface{})
	if ok {
		delete(mapc, "sort")
		delete(mapc, "page")
		delete(mapc, "size")

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

// Get perform a find one.
func (c *Collection) Get(instance Model) error {
	result := c.db.First(instance)
	return result.Error
}

// Find perform a find all.
func (c *Collection) Find(instance interface{}) error {
	result := c.db.Find(instance)
	return result.Error
}

// Update perform an update. You must perform a Conditions first in order to use the Update method.
func (c *Collection) Update(instance Model) error {
	result := c.db.Save(instance)
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

// Delete perform a delete.
func (c *Collection) Delete(instance Model) error {
	result := c.db.Delete(instance)
	return result.Error
}

func (c *Collection) Exists(instance Model) (bool, error) {
	result := c.db.First(instance)
	if result.Error != nil {
		if errors.Is(result.Error, ErrRecordNotFound) {
			return false, nil
		}

		return false, result.Error
	}
	return true, nil
}
