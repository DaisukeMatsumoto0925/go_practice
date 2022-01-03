package interfaces

import "github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/repositories"

type SQLHandler interface {
	Exec(string, ...interface{}) SQLHandler
	Find(interface{}, ...interface{}) SQLHandler
	First(interface{}, ...interface{}) SQLHandler
	Raw(string, ...interface{}) SQLHandler
	Create(interface{}) SQLHandler
	Save(interface{}) SQLHandler
	Delete(interface{}) SQLHandler
	Where(interface{}, ...interface{}) SQLHandler
	Related(interface{}, ...string) SQLHandler
	Preload(column string, conditions ...interface{}) SQLHandler
	Model(interface{}) SQLHandler
	Table(string) SQLHandler
	Joins(string, ...interface{}) SQLHandler
	Select(string, ...interface{}) SQLHandler
	TransactAndReturnData(func(repositories.Transaction) (interface{}, error)) (interface{}, error)
	GetError() error
	RecordNotFound() bool
	Close() error
	FromTransaction(repositories.Transaction) SQLHandler
	Transactionable()
}
