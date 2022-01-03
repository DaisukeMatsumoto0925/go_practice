package db

import (
	"fmt"
	"log"
	"os"

	"github.com/DaisukeMatsumoto0925/oapi_gen/config"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/repositories"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/interfaces"
	"github.com/jinzhu/gorm"
)

// A SQLHandler belong to the infrastructure layer.
type SQLHandler struct {
	Conn *gorm.DB
}

func NewSQLHandler() *SQLHandler {
	connectionString := genConnectString()
	conn, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("db connection error: ", err)
	}

	// 開発時のみSQLのデバッグログを表示する
	if os.Getenv("GO_ENV") == "development" {
		conn.LogMode(true) // debug mode
	}

	return &SQLHandler{
		Conn: conn,
	}
}

func genConnectString() string {
	conf := config.Get()
	USER := conf.Db.User
	PASSWORD := conf.Db.Password
	PROTOCOL := fmt.Sprintf("tcp(%s:%s)", conf.Db.Host, conf.Db.Port)
	DBNAME := conf.Db.Name
	QUERY := "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"

	return USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME + QUERY
}

func (s *SQLHandler) Find(out interface{}, where ...interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Find(out, where...)}
}

func (s *SQLHandler) Exec(sql string, values ...interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Exec(sql, values...)}
}

func (s *SQLHandler) First(out interface{}, where ...interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.First(out, where...)}
}

func (s *SQLHandler) Raw(sql string, values ...interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Raw(sql, values...)}
}

func (s *SQLHandler) Create(value interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Create(value)}
}

func (s *SQLHandler) Save(value interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Save(value)}
}

func (s *SQLHandler) Delete(value interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Delete(value)}
}

func (s *SQLHandler) Where(query interface{}, args ...interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Where(query, args...)}
}

func (s *SQLHandler) Related(value interface{}, foreignKeys ...string) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Related(value, foreignKeys...)}
}

func (s *SQLHandler) Preload(column string, conditions ...interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Preload(column, conditions...)}
}

func (s *SQLHandler) Model(value interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Model(value)}
}

func (s *SQLHandler) Table(name string) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Table(name)}
}

func (s *SQLHandler) Joins(query string, args ...interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Joins(query, args...)}
}

func (s *SQLHandler) Select(query string, args ...interface{}) interfaces.SQLHandler {
	return &SQLHandler{Conn: s.Conn.Select(query, args...)}
}

func (s *SQLHandler) TransactAndReturnData(txFunc func(repositories.Transaction) (interface{}, error)) (data interface{}, err error) {
	tx := s.Conn.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	data, err = txFunc(&SQLHandler{Conn: tx})
	return
}

func (s *SQLHandler) GetError() error {
	return s.Conn.Error
}

func (s *SQLHandler) RecordNotFound() bool {
	return s.Conn.RecordNotFound()
}

func (s *SQLHandler) Close() error {
	return s.Conn.Close()
}

func (s *SQLHandler) FromTransaction(tx repositories.Transaction) interfaces.SQLHandler {
	sqlhandler, ok := tx.(*SQLHandler)
	if ok {
		return sqlhandler
	}

	return s
}

func (s *SQLHandler) Transactionable() {
}
