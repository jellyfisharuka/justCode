package entity3

import (
    "database/sql"
    "errors"
)

type Entity3Repository interface {
    GetAll() ([]Entity3, error)
    GetById(id int) (Entity3, error)
    Create(entity Entity3) (int, error)
    Update(id int, entity Entity3) error
    Delete(id int) error
}

type Entity3RepositoryImpl struct {
    db *sql.DB
}

func NewEntity3Repository(db *sql.DB) Entity3Repository {
    return &Entity3RepositoryImpl{db}
}

func (repo *Entity3RepositoryImpl) GetAll() ([]Entity3, error) {
    rows, err := repo.db.Query("SELECT * FROM entity3")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var entities []Entity3
    for rows.Next() {
        var entity Entity3
        if err := rows.Scan(&entity.ID, &entity.Description); err != nil {
            return nil, err
        }
        entities = append(entities, entity)
    }

    return entities, nil
}

func (repo *Entity3RepositoryImpl) GetById(id int) (Entity3, error) {
    var entity Entity3
    err := repo.db.QueryRow("SELECT * FROM entity3 WHERE ID = $1", id).Scan(&entity.ID, &entity.Description)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return Entity3{}, errors.New("Entity3 not found")
        }
        return Entity3{}, err
    }

    return entity, nil
}

func (repo *Entity3RepositoryImpl) Create(entity Entity3) (int, error) {
    var id int
    err := repo.db.QueryRow("INSERT INTO entity3(description) VALUES($1) RETURNING id", entity.Description).Scan(&id)
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (repo *Entity3RepositoryImpl) Update(id int, entity Entity3) error {
    _, err := repo.db.Exec("UPDATE entity3 SET description = $1 WHERE id = $2", entity.Description, id)
    return err
}

func (repo *Entity3RepositoryImpl) Delete(id int) error {
    _, err := repo.db.Exec("DELETE FROM entity3 WHERE id = $1", id)
    return err
}
