package scoring

import (
	"errors"

	"github.com/gocql/gocql"
)

type DB interface {
	Connect() error
	Disconnect()
	InitTables() error
	IncrementBackLinksCount() error
}

type Cassandra struct {
	session      *gocql.Session
	init         bool
	keyspaceName string
}

// Connect Connects to Multiple Clusters of Apache Cassandra
func (t *Cassandra) Connect(hosts []string, keyspaceName string) (err error) {
	cluster := gocql.NewCluster(hosts...)
	t.keyspaceName = keyspaceName
	cluster.Keyspace = keyspaceName
	t.session, err = cluster.CreateSession()
	return
}

// Disconnect Disconnects the client
func (t *Cassandra) Disconnect() {
	t.session.Close()
}

// InitTables Initialize the Tables
func (t *Cassandra) InitTables() (err error) {
	if t.init {
		return errors.New("table is already initialized")
	}

	// Creating Table `backlinks`
	err = t.session.Query(`
		CREATE TABLE IF NOT EXISTS backlinks (
			url TEXT PRIMARY KEY, 
			backlink_count COUNTER
		);
	`).Exec()
	if err != nil {
		return errors.New("unable to create table backlinks")
	}

	t.init = true
	return
}
