package mocks

import (
	"database/sql"
)

type MockDB struct {
	mock *sql.DB
}

func NewMockDB(mock *sql.DB) *MockDB {
	return &MockDB{mock: mock}
}

func (m *MockDB) GetConnection() *sql.DB {
	return m.mock
}
