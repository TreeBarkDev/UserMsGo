package service

import (
    "context"

    "github.com/gocql/gocql"
    "go-cassandra-demo-service/internal/model"
)

func InsertUser(session *gocql.Session, u *model.User) error {
    return session.Query(
        "INSERT INTO users (id, name, email) VALUES (?, ?, ?)",
        u.ID, u.Name, u.Email,
    ).WithContext(context.Background()).Exec()
}
