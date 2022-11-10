package orm

import (
	"context"

	"github.com/tiger-guo/benchmark/mysql/types"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
)

// NewORM new object relation map.
func NewORM() (*ORM, error) {
	client, err := sqlx.Open("mysql", types.MysqlSource)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(); err != nil {
		return nil, err
	}

	client.SetMaxIdleConns(types.MysqlMaxIdle)
	client.SetMaxOpenConns(types.MysqlMaxConn)

	return &ORM{
		db: client,
	}, nil
}

// ORM defines data operation struct.
type ORM struct {
	db *sqlx.DB
}

// Close orm.
func (orm *ORM) Close() {
	orm.db.Close()
	return
}

// Get one data and decode into dest *struct{}.
func (orm *ORM) Get(ctx context.Context, dest interface{}, expr string, args ...interface{}) error {
	err := orm.db.GetContext(ctx, dest, expr, args...)
	if err != nil {
		return err
	}

	return nil
}

// Select a collection of data, and decode into dest *[]struct{}.
func (orm *ORM) Select(ctx context.Context, dest interface{}, expr string, args ...interface{}) error {
	expr, args, err := sqlx.In(expr, args...)
	if err != nil {
		return err
	}

	err = orm.db.SelectContext(ctx, dest, expr, args...)
	if err != nil {
		return err
	}

	return nil
}

// Insert a row data to db
func (orm *ORM) Insert(ctx context.Context, expr string, data interface{}) error {
	_, err := orm.db.NamedExecContext(ctx, expr, data)
	if err != nil {
		return err
	}

	return err
}
