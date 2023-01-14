package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/transferMVP/transfer.webapp/internal/config"
	"github.com/transferMVP/transfer.webapp/internal/models"
	"time"
)

var db *pgxpool.Pool

func InitPool() error {
	var databaseUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Config.Pg.User, config.Config.Pg.Pass, config.Config.Pg.Host, config.Config.Pg.Port, config.Config.Pg.DbName)

	poolConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		return err
	}
	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return err
	}
	return nil
}

func AddUser(user *models.User) (string, error) {
	id := ""
	if err := db.QueryRow(context.Background(), `insert into users
			(first_name, last_name, email, phone, сountry_code, birthdate, password, created_at )
			VALUES ($1, $2, $3, $4,$5, $6, $7, $8) returning (id)`,
		user.FirstName, user.LastName, user.Email,
		user.Phone, user.CountryCode, user.BirthDate,
		user.Password, time.Now().UTC()).Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func GetUser(id string) (models.User, error) {
	fmt.Println(id)
	var user models.User
	rows, err := db.Query(context.Background(), `select  "first_name", "last_name", "email", 
	"phone" , "blocked", "created_at", "updated_at" from users where id=$1`,
		id)

	if err != nil {
		return user, err
	}
	for rows.Next() {
		var fn string
		var ln string
		var em string
		var ph string
		var bl bool
		var cr time.Time
		var up time.Time
		err := rows.Scan(&fn, &ln, &em, &ph, &bl, &cr, &up)
		if err != nil {
			return models.User{}, err
		}
		user.Id = id
		user.FirstName = fn
		user.LastName = ln
		user.Email = em
		user.Phone = ph
		user.Blocked = bl
		user.CreatedAt = cr
		user.UpdatedAt = up
	}
	return user, nil
}
