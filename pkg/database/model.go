package database

import (
	"Walker/pkg/contract"
	"Walker/pkg/database/drivers"
	"database/sql"
	"gorm.io/gorm"
)

type Model struct {
	Db *gorm.DB
}

func (m *Model) DBConnector(config contract.Config) contract.Model {
	switch config.GetString("db.connection") {
	case "mysql":
		m.Db = (&drivers.Mysql{}).MysqlConnector(config)
	}
	return m
}

func (m *Model) Model(value interface{}) contract.Model {
	m.Db = m.Db.Model(value)
	return m
}
func (m *Model) Where(query interface{}, args ...interface{}) contract.Model {
	m.Db = m.Db.Where(query, args...)
	return m
}

func (m *Model) Find(dest interface{}, conds ...interface{}) contract.Model {
	m.Db = m.Db.Find(dest, conds...)
	return m
}
func (m *Model) First(dest interface{}, conds ...interface{}) contract.Model {
	m.Db = m.Db.First(dest, conds...)
	return m
}
func (m *Model) Not(query interface{}, args ...interface{}) contract.Model {
	m.Db = m.Db.Not(query, args...)
	return m
}
func (m *Model) Or(query interface{}, args ...interface{}) contract.Model {
	m.Db = m.Db.Or(query, args...)
	return m
}
func (m *Model) Select(query interface{}, args ...interface{}) contract.Model {
	m.Db = m.Db.Select(query, args...)
	return m
}
func (m *Model) Rows() (*sql.Rows, error) {
	Rows, err := m.Db.Rows()
	return Rows, err
}
func (m *Model) Order(value interface{}) contract.Model {
	m.Db = m.Db.Order(value)
	return m
}
func (m *Model) Offset(offset int) contract.Model {
	m.Db = m.Db.Offset(offset)
	return m
}
func (m *Model) Limit(limit int) contract.Model {
	m.Db = m.Db.Limit(limit)
	return m
}
func (m *Model) Group(name string) contract.Model {
	m.Db = m.Db.Group(name)
	return m
}
func (m *Model) Having(query interface{}, args ...interface{}) contract.Model {
	m.Db = m.Db.Having(query, args...)
	return m
}
func (m *Model) Distinct(args ...interface{}) contract.Model {
	m.Db = m.Db.Distinct(args...)
	return m
}
func (m *Model) Debug() contract.Model {
	m.Db = m.Db.Debug()
	return m
}
