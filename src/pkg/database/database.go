package database

import "errors"

type IDataBase interface {
	Put(string, string) (bool, error)
	Remove(string) (bool, error)
	Read(string) (string, error)
	Close() error
}

var db IDataBase

func GetDataBase() (IDataBase, error) {
	if db == nil {
		return nil, errors.New("database wasn't initialized")
	}
	return db, nil
}

func InitDataBase(url string) error {
	var err error

	if db == nil {
		if url == "test" {
			db, err = makeDbMock()
		} else {
			db, err = makeMongoDB(url)
		}
	} else {
		err = errors.New("database is already initialized")
	}

	return err
}

func DisposeDataBase() error {
	var err error
	if db != nil {
		err = db.Close()
		db = nil
	}
	return err
}
