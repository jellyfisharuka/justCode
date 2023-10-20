package entity1

import (
    "database/sql"
    "errors"
)
type Entity1Repository interface {
    GetAll() ([]Entity1, error)
    GetById(id int) (Entity1, error)
    Create(entity Entity1) (int, error)
    Update(id int, entity Entity1) error
    Delete(id int) error
}

type Entity1RepositoryImpl struct {
    db *sql.DB
}

func NewEntity1Repository(db *sql.DB) Entity1Repository {
    return &Entity1RepositoryImpl{db}
}

func (repo *Entity1RepositoryImpl) GetAll() ([]Entity1, error) {
    rows, err := repo.db.Query("SELECT * FROM entity1")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var entities []Entity1
    for rows.Next() {
        var entity Entity1
        if err := rows.Scan(&entity.ID, &entity.Name); err != nil {
            return nil, err
        }
        entities = append(entities, entity)
    }

    return entities, nil
}

func (repo *Entity1RepositoryImpl) GetById(id int) (Entity1, error) {
    var entity Entity1
    err := repo.db.QueryRow("SELECT * FROM entity1 WHERE ID = $1", id).Scan(&entity.ID, &entity.Name)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return Entity1{}, errors.New("Entity1 not found")
        }
        return Entity1{}, err
    }

    return entity, nil
}

func (repo *Entity1RepositoryImpl) Create(entity Entity1) (int, error) {
    var id int
    err := repo.db.QueryRow("INSERT INTO entity1(name) VALUES($1) RETURNING id", entity.Name).Scan(&id)
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (repo *Entity1RepositoryImpl) Update(id int, entity Entity1) error {

    _, err := repo.db.Exec("UPDATE entity1 SET name = $1 WHERE id = $2", entity.Name, id)
    return err
}

func (repo *Entity1RepositoryImpl) Delete(id int) error {

    _, err := repo.db.Exec("DELETE FROM entity1 WHERE id = $1", id)
    return err
}
