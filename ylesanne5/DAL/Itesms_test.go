package DAL_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/keijoraamat/IDU1550/ylesanne5/DAL"
	"github.com/keijoraamat/IDU1550/ylesanne5/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.NewWithDSN("file::memory:?cache=shared")
	if err != nil {
		panic("failed to open mock sql DB")
	}

	defer db.Close()

	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to open gorm DB")
	}

	err = gormDB.AutoMigrate(&models.Item{})
	if err != nil {
		panic("failed to create table")
	}

	return gormDB, mock
}

func teardownMockDB(db *gorm.DB) {
	// Clear the database
	db.Exec("DELETE FROM items")

}

func TestCreateItem(t *testing.T) {
	db, mock := setupMockDB()
	repo := DAL.ItemRepository{DB: db}

	item := models.Item{Name: "Test Item", Price: 10}

	// Expectations
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO items").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	id := repo.CreateItem(&item)

	if id == 0 {
		t.Errorf("CreateItem was incorrect, got: %d, want: %s.", id, "greater than 0")
	}

	teardownMockDB(db)
}

func TestGetAllItems(t *testing.T) {
	db, _ := setupMockDB()
	repo := DAL.ItemRepository{DB: db}

	// Mock data
	expectedItems := []models.Item{
		{Name: "Item1", Price: 10},
		{Name: "Item2", Price: 20},
	}

	// Insert mock data into the mock DB
	for _, item := range expectedItems {
		repo.CreateItem(&item)
	}

	// Call the function to test
	items := repo.GetAllItems()

	// Check that the function returns the correct data
	if len(items) != len(expectedItems) {
		t.Errorf("expected %d items, got %d", len(expectedItems), len(items))
	}

	for i, item := range items {
		if item.Price != expectedItems[i].Price || item.Name != expectedItems[i].Name {
			t.Errorf("expected item %v, got %v", expectedItems[i], item)
		}
	}

	teardownMockDB(db)
}
