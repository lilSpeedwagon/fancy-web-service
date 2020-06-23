package database

import "errors"

type dbMock struct {
	data map[string]string
}

func (db dbMock) Put(key, value string) error {
	db.data[key] = value
	return nil
}

func (db dbMock) Remove(key string) error {
	delete(db.data, key)
	return nil
}

func (db dbMock) Read(key string) (string, error) {
	val, isExist := db.data[key]
	if !isExist {
		return "", errors.New("key is not exist")
	}
	return val, nil
}

func makeDbMock() IDataBase {
	db := dbMock{}
	db.data = make(map[string]string)
	return db
}
