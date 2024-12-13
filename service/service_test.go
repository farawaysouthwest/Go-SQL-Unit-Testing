package service

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	mock_database "testingExample/mocks"
)

type testContext struct {
	ctx          context.Context
	mockCtrl     *gomock.Controller
	mockDatabase *mock_database.MockDatabase
	gorm         *gorm.DB
	sqlCtrl      sqlmock.Sqlmock
}

func newTestContext(t *testing.T) *testContext {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	mockDatabase := mock_database.NewMockDatabase(mockCtrl)

	db, sqlCtrl, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
		Conn:                      db,
	})

	g, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return &testContext{
		ctx:          ctx,
		mockCtrl:     mockCtrl,
		mockDatabase: mockDatabase,
		gorm:         g,
		sqlCtrl:      sqlCtrl,
	}
}

func TestService_GetAll(t *testing.T) {

	// Create a new test context
	tc := newTestContext(t)
	defer tc.mockCtrl.Finish()

	// setup mock expectations
	userName := "John"

	tc.mockDatabase.EXPECT().GetConnection().Return(tc.gorm)

	tc.sqlCtrl.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_models`")).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(userName))

	// Create a new service
	s := NewService(tc.mockDatabase)

	users, err := s.GetAll(tc.ctx)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when getting all users", err)
	}

	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}

	if users[0].Name != userName {
		t.Fatalf("expected user name to be 'John', got %s", users[0].Name)
	}
}
