package entity2

import (
	"database/sql"
	"errors"
)

type Entity2Repository interface {
	GetAll() ([]Entity2, error)
	GetById(id int) (Entity2, error)
	Create(entity Entity2) (int, error)
	Update(id int, entity Entity2) error
	Delete(id int) error
}

type Entity2RepositoryImpl struct {
	db *sql.DB
}

func NewEntity2Repository(db *sql.DB) Entity2Repository {
	return &Entity2RepositoryImpl{db}
}

func (repo *Entity2RepositoryImpl) GetAll() ([]Entity2, error) {
	rows, err := repo.db.Query("SELECT * FROM entity2")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []Entity2
	for rows.Next() {
		var entity Entity2
		if err := rows.Scan(&entity.ID, &entity.Name); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (repo *Entity2RepositoryImpl) GetById(id int) (Entity2, error) {
	var entity Entity2
	err := repo.db.QueryRow("SELECT * FROM entity2 WHERE ID = $1", id).Scan(&entity.ID, &entity.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Entity2{}, errors.New("Entity2 not found")
		}
		return Entity2{}, err
	}

	return entity, nil
}

func (repo *Entity2RepositoryImpl) Create(entity Entity2) (int, error) {
	var id int
	err := repo.db.QueryRow("INSERT INTO entity2(name) VALUES($1) RETURNING id", entity.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *Entity2RepositoryImpl) Update(id int, entity Entity2) error {
	_, err := repo.db.Exec("UPDATE entity2 SET name = $1 WHERE id = $2", entity.Name, id)
	return err
}

func (repo *Entity2RepositoryImpl) Delete(id int) error {
	_, err := repo.db.Exec("DELETE FROM entity2 WHERE id = $1", id)
	return err
}
