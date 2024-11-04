package repository

import (
	"testing"

	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestAuthPostgres_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewUserPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		input   domain.User
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO users").WithArgs("test", "test", "test").WillReturnRows(rows)
			},
			input: domain.User{

				Login:    "test",
				Password: "test",
			},
			want: 1,
		},
		{
			name: "Empty fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").WithArgs("test", "test", "").WillReturnRows(rows)
			},
			input: domain.User{

				Login:    "test",
				Password: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.CreateUser(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func testAuthPostgres_getUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	r := NewUserPostgres(db)

	type args struct {
		username string
		password string
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		want    domain.User
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "username", "password"}).AddRow(1, "test", "test", "test")
				mock.ExpectQuery("SELECT (.+) FROM users").WithArgs("test", "password").WillReturnRows(rows)
			},
			input: args{"test", "test"},
			want:  domain.User{1, "test", "test"},
		},
		{
			name: "Now found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "username", "password"})
				mock.ExpectQuery("SELECT (.+) FROM users").WithArgs("not", "found").WillReturnRows(rows)
			},
			input:   args{"not", "found"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := r.GetUser(tt.input.username, tt.input.password)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
