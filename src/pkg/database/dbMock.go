package database

type dbMock struct {
	data map[string]string
}

func (db dbMock) Put(key, value string) (bool, error) {
	_, hasKey := db.data[key]
	db.data[key] = value
	isInserted := !hasKey
	return isInserted, nil
}

func (db dbMock) Remove(key string) (bool, error) {
	_, hasKey := db.data[key]
	delete(db.data, key)
	return hasKey, nil
}

func (db dbMock) Read(key string) (string, error) {
	val, isExist := db.data[key]
	if !isExist {
		return "", nil
	}
	return val, nil
}

func (db dbMock) Close() error {
	return nil
}

func makeDbMock() (IDataBase, error) {
	db := dbMock{}
	db.data = make(map[string]string)
	return db, nil
}
