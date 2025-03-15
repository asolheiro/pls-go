package task

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Tasks struct {
	Tasks []Task
}

type Task struct {
	ID int64
	Name string
	Done bool
}

type TasksManager struct {
	db *sql.DB
}

func New(dbPath string) (*TasksManager, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	tm := &TasksManager{db: db}
	if err := tm.initDB(); err != nil {
		db.Close()
		return nil, err
	}

	return tm, nil
}

func (tm *TasksManager) Close() error {
	return tm.db.Close()
}

func (tm *TasksManager) initDB() error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := tm.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func (tm *TasksManager) Add(name string) (*Task, error) {
	if name == "" {
		return nil, errors.New("task name cannot be empty")
	}

	query := "INSERT INTO todos (name, done) VALUES (?, 0)"
	result, err := tm.db.Exec(query, name)
	if err != nil {
		return nil, fmt.Errorf("failed to add todo: %w", err)
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	return &Task{
		ID:   ID,
		Name: name,
		Done: false,
	}, nil
}

func (tm *TasksManager) Get(ID int64) (*Task, error) {
	query := "SELECT ID, name, done FROM todos WHERE ID = ?"
	row := tm.db.QueryRow(query, ID)

	var todo Task
	err := row.Scan(&todo.ID, &todo.Name, &todo.Done)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("todo with ID %d not found", ID)
		}
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	return &todo, nil
}

func (tm *TasksManager) GetAll() ([]*Task, error) {
	query := "SELECT ID, name, done FROM todos ORDER BY ID"
	rows, err := tm.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query todos: %w", err)
	}
	defer rows.Close()

	var todos []*Task
	for rows.Next() {
		var todo Task
		if err := rows.Scan(&todo.ID, &todo.Name, &todo.Done); err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, &todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return todos, nil
}

func (tm *TasksManager) Update(todo *Task) error {
	if todo == nil {
		return errors.New("todo cannot be nil")
	}

	query := "UPDATE todos SET name = ?, done = ? WHERE id = ?"
	result, err := tm.db.Exec(query, todo.Name, todo.Done, todo.ID)
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %d not found", todo.ID)
	}

	return nil
}

func (tm *TasksManager) MarkDone(id int64) error {
	query := "UPDATE todos SET done = 1 WHERE id = ?"
	result, err := tm.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to mark todo as done: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %d not found", id)
	}

	return nil
}

func (tm *TasksManager) MarkUndone(id int64) error {
	query := "UPDATE todos SET done = 0 WHERE id = ?"
	result, err := tm.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to mark todo as undone: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %d not found", id)
	}

	return nil
}

func (tm *TasksManager) Delete(id int64) error {
	query := "DELETE FROM todos WHERE id = ?"
	result, err := tm.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %d not found", id)
	}

	return nil
}

func (tm *TasksManager) DeleteAll() error {
	query := "DELETE FROM todos"
	_, err := tm.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to delete all todos: %w", err)
	}

	return nil
}
