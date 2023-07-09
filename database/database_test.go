package database

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	var DB Database
	err := DB.Init("./db/testdb.db")

	if err != nil {
		t.Fatalf("didn't Initialize the db got error '%v' ", err)
	}
}

func TestCreatingAModel(t *testing.T) {
	var DB Database
	DB.Init("./db/testdb.db")

	type TestModel struct {
		ID         uint `gorm:"primaryKey"`
		TestField1 string
		TestField2 string
	}

	err := DB.CreateTable(TestModel{})
	if err != nil {
		t.Fatalf("didn't create a table , got error : '%v'", err)
	}
}

func TestAdd(t *testing.T) {
	var DB Database
	DB.Init("./db/testdb.db")

	type TestModel struct {
		ID         uint `gorm:"primaryKey"`
		TestField1 string
		TestField2 string
	}

	result := DB.Add(&TestModel{TestField1: "testValue1", TestField2: "testValue2"})

	err := result.Error
	if err != nil {
		t.Fatalf("didn't create a recorde ,got erorr: '%v'", err)
	}
}

func TestFetch(t *testing.T) {
	var DB Database
	DB.Init("./db/testdb.db")

	type TestModel struct {
		ID         uint `gorm:"primaryKey"`
		TestField1 string
		TestField2 string
	}

	var tst1 TestModel
	rs := DB.FindOneByID(&tst1, 1)
	if rs.Error != nil {
		t.Fatalf("cannot get the Item by Id , got err: '%v'", rs.Error)
	}
	fmt.Println(tst1)

	var tst2 TestModel
	rs2 := DB.FindOneByCol(&tst2, "test_field1", "testValue1")
	if rs2.Error != nil {
		t.Fatalf("cannot get the Item by Id , got err: '%v'", rs2.Error)
	}
	fmt.Println(tst2)
}
