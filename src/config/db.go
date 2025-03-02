package config

import (
	"log"
	"os"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

var DBSession *gocql.Session

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOST"))
	cluster.Port = 9042
	cluster.Keyspace = os.Getenv("CASSANDRA_KEYSPACE")
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}

	DBSession = session
	log.Println("Connected to Cassandra successfully")
}

func CloseDB() {
	if DBSession != nil {
		DBSession.Close()
		log.Println("Database connection closed")
	}
}
