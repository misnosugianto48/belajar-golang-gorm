package belajargolanggorm

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=belajar_golang_gormdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestMigrator(t *testing.T) {
	err := db.Migrator().AutoMigrate(&GuestBook{})
	assert.Nil(t, err)
}

// gunakan migrator untuk pengetesan, disarankan menggunakan database migration

func TestExecuteSQL(t *testing.T) {
	t.Run("Exec SQL", func(t *testing.T) {
		err := db.Exec("insert into sample(id, name) values(?,?)", "1", "Misno").Error
		assert.Nil(t, err)

		err = db.Exec("insert into sample(id, name) values(?,?)", "2", "Sugianto").Error
		assert.Nil(t, err)

		err = db.Exec("insert into sample(id, name) values(?,?)", "3", "Anto").Error
		assert.Nil(t, err)

		err = db.Exec("insert into sample(id, name) values(?,?)", "4", "Knock").Error
		assert.Nil(t, err)
	})

	t.Run("RawSQL", func(t *testing.T) {
		var sample Sample
		err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error
		assert.Nil(t, err)
		assert.Equal(t, "Misno", sample.Name)
	})

	t.Run("RawsSQL", func(t *testing.T) {
		var samples []Sample
		err := db.Raw("select id, name from sample").Scan(&samples).Error

		assert.Nil(t, err)
		assert.Equal(t, 4, len(samples))
	})

	t.Run("ScanRows", func(t *testing.T) {
		rows, err := db.Raw("select id, name from sample").Rows()
		assert.Nil(t, err)
		defer rows.Close()

		var samples []Sample
		for rows.Next() {
			db.ScanRows(rows, &samples)
		}
		assert.Equal(t, 4, len(samples))
	})
}

func TestCRUDInterface(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		user := User{
			ID: "1",
			Name: Name{
				FirstName: "Misno",
				LastName:  "Sugianto",
			},
			Password: "secret_password",
		}

		res := db.Create(&user)
		assert.Nil(t, res.Error)
		assert.Equal(t, int64(1), res.RowsAffected)
	})
	t.Run("Batch Insert", func(t *testing.T) {
		var users []User
		for i := 2; i < 10; i++ {
			users =
				append(users, User{
					ID:       strconv.Itoa(i),
					Password: "secret",
					Name: Name{
						FirstName: "User " + strconv.Itoa(i),
					},
				})
		}

		res := db.Create(&users)
		assert.Nil(t, res.Error)
		assert.Equal(t, int64(8), res.RowsAffected)
	})
}
