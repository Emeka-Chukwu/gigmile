package data

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"
)

const dbTimeout = time.Second * 5

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Country: Country{},
	}
}

type Models struct {
	Country Country
}

type Country struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	ShortName      string    `json:"shortname"`
	Continent      string    `json:"continent,"`
	Is_Operational bool      `json:"is_operational"`
	Created_At     time.Time `json:"created_at"`
	Updated_At     time.Time `json:"updated_at"`
}

func (c *Country) GetAll() ([]*Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, shortname, continent, is_operational, created_at, updated_at
	from countries order by created_at desc`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {

		return []*Country{}, err
	}
	defer rows.Close()

	var countries []*Country

	for rows.Next() {
		var country Country
		err := rows.Scan(
			&country.ID,
			&country.Name,
			&country.ShortName,
			&country.Continent,
			&country.Is_Operational,
			&country.Created_At,
			&country.Updated_At,
		)
		if err != nil {
			return []*Country{}, err
		}

		countries = append(countries, &country)
	}

	if countries == nil {
		countries = []*Country{}
	}
	return countries, nil
}

func (c *Country) GetOne(id int) (*Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, shortname, continent, is_operational, created_at, updated_at
	from countries where id = $1`

	var country Country
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&country.ID,
		&country.Name,
		&country.ShortName,
		&country.Continent,
		&country.Is_Operational,
		&country.Created_At,
		&country.Updated_At,
	)

	if err != nil {
		return nil, err
	}

	return &country, nil
}

func (c *Country) GetCountry(name, short string) (*Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, shortname, continent, is_operational, created_at, updated_at from countries where name = $1 or shortname = $2`

	var country Country
	row := db.QueryRowContext(ctx, query, name, short)
	err := row.Scan(
		&country.ID,
		&country.Name,
		&country.ShortName,
		&country.Continent,
		&country.Is_Operational,
		&country.Created_At,
		&country.Updated_At,
	)

	if err != nil {
		return nil, err
	}

	return &country, nil
}

func (country *Country) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	update := `update countries set
		name = $1,
		shortname = $2,
		continent = $3,
		is_operational = $4,
		updated_at = $5
		where id = $6
	`

	_, err := db.ExecContext(ctx, update,

		&country.Name,
		&country.ShortName,
		&country.Continent,
		&country.Is_Operational,
		time.Now(),
		&country.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *Country) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from countries where id = $1`

	value, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	index, _ := value.RowsAffected()
	if index == 0 {
		return errors.New("no record found")
	}

	return nil
}

func (country *Country) InsertCountry() (*Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `insert into countries (name, shortname, continent, is_operational, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6) returning id, name, shortname, continent, is_operational, created_at, updated_at`

	err := db.QueryRowContext(ctx, stmt,
		strings.ToLower(country.Name),
		strings.ToLower(country.ShortName),
		country.Continent,
		country.Is_Operational,
		time.Now(),
		time.Now(),
	).Scan(
		&country.ID,
		&country.Name,
		&country.ShortName,
		&country.Continent,
		&country.Is_Operational,
		&country.Created_At,
		&country.Updated_At,
	)

	if err != nil {
		return &Country{}, err
	}

	return country, nil
}
