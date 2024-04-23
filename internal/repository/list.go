package repository

import (
	"fmt"
	"strings"
	"todoApp/internal/models"

	"github.com/jmoiron/sqlx"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{
		db: db,
	}
}

func (r *ListPostgres) Create(userID int, list models.List) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *ListPostgres) GetAll(userID int) ([]models.List, error) {
	var lists []models.List
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userID)
	return lists, err
}

func (r *ListPostgres) GetByID(userID, listID int) (models.List, error) {
	var list models.List
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userID, listID)
	return list, err
}

func (r *ListPostgres) Update(userId, listId int, input models.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ListPostgres) Delete(userID, listID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteUsersListQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND list_id=$2", usersListsTable)
	_, err = tx.Exec(deleteUsersListQuery, userID, listID)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteListQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", todoListsTable)
	_, err = tx.Exec(deleteListQuery, listID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
