package contract

import "database/sql"

type Model interface {
	Model(value interface{}) Model
	Where(query interface{}, args ...interface{}) Model
	Find(dest interface{}, conds ...interface{}) Model
	First(dest interface{}, conds ...interface{}) Model
	Not(query interface{}, args ...interface{}) Model
	Or(query interface{}, args ...interface{}) Model
	Select(query interface{}, args ...interface{}) Model
	Rows() (*sql.Rows, error)
	Debug() Model
	Order(value interface{}) Model
	Offset(offset int) Model
	Limit(limit int) Model
	Group(name string) Model
	Having(query interface{}, args ...interface{}) Model
	Distinct(args ...interface{}) Model
}
