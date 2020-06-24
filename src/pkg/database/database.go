package database

type IDataBase interface {
	Put(string, string) (bool, error)
	Remove(string) (bool, error)
	Read(string) (string, error)
}

var db IDataBase

func OpenDataBase(url string) (IDataBase, error) {
	if db == nil {
		db = makeDbMock()
	}
	return db, nil
}
