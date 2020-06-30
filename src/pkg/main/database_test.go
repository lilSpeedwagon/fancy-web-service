package main

import (
	"pkg/database"
	"testing"
)

const (
	dbUrl = "test"
)

func openDB(t *testing.T) database.IDataBase {
	if err := database.InitDataBase(dbUrl); err != nil {
		t.Errorf("Cannot init databse with url: %s.", dbUrl)
	}

	db, err := database.GetDataBase()
	if db == nil {
		t.Errorf("Database object is nil.")
	}
	if err != nil {
		t.Errorf("Error occured while getting database: %s.", err.Error())
	}
	return db
}

func closeDB(t *testing.T) {
	err := database.DisposeDataBase()
	if err != nil {
		t.Errorf("Error occured while disposing database: %s.", err.Error())
	}
}

func TestDatabaseConnection(t *testing.T) {
	openDB(t)
	closeDB(t)
}

//noinspection GoBoolExpressions
func TestDataBaseCrud(t *testing.T) {
	db := openDB(t)

	key := "kkk"
	value := "vvv"

	// put
	expectedPut := true

	isInserted, errPut := db.Put(key, value)
	if errPut != nil {
		t.Errorf("IDataBase.Put returned unexpected error: %s.", errPut.Error())
	}
	if isInserted != expectedPut {
		t.Errorf("IDataBase.Put returned: %t. Expected: %t.", isInserted, expectedPut)
	}

	// read
	expectedRead := value

	readVal, errRead := db.Read(key)
	if errRead != nil {
		t.Errorf("IDataBase.Read returned unexpected error: %s.", errRead.Error())
	}
	if readVal != expectedRead {
		t.Errorf("IDataBase.Read returned: %s. Expected: %s.", readVal, expectedRead)
	}

	// remove
	expectedRemove := true

	isRemoved, errRemove := db.Remove(key)
	if errRemove != nil {
		t.Errorf("IDataBase.Remove returned unexpected error: %s.", errRemove.Error())
	}
	if isRemoved != expectedRemove {
		t.Errorf("IDataBase.Remove returned: %t. Expected: %t.", isRemoved, expectedRemove)
	}

	// read2
	expectedRead2 := ""

	readVal2, errRead2 := db.Read(key)
	if errRead2 != nil {
		t.Errorf("IDataBase.Read returned unexpected error: %s.", errRead2.Error())
	}
	if readVal2 != expectedRead2 {
		t.Errorf("IDataBase.Read returned: %s. Expected: %s.", readVal2, expectedRead2)
	}

	closeDB(t)
}
