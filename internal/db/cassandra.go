package db

import (
    "log"
    "time"

    "github.com/gocql/gocql"
    "go-cassandra-demo-service/internal/config"
)

func Connect(cfg config.Config) *gocql.Session {
    cluster := gocql.NewCluster(cfg.CassandraHost)
    cluster.Keyspace = cfg.Keyspace
    cluster.Consistency = gocql.Quorum
    cluster.Timeout = 10 * time.Second
    session, err := cluster.CreateSession()
    if err != nil {
        log.Fatalf("Unable to connect to Cassandra: %v", err)
    }
    return session
}
