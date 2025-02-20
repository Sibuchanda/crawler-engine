package scoring

import (
	"errors"
	"time"

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
	err = t.session.Query(`CREATE TABLE IF NOT EXISTS backlinks (
			url TEXT PRIMARY KEY, 
			backlink_count COUNTER
		);`).Exec()
	if err != nil {
		return errors.New("unable to create table backlinks")
	}

	// Creating Table `last_modified`
	err = t.session.Query(`CREATE TABLE IF NOT EXISTS last_modified (
			url TEXT PRIMARY KEY,
			last_modified_time TIMESTAMP
		);`).Exec()
	if err != nil {
		return errors.New("unable to create table last_modified")
	}

	t.init = true
	return
}

// UpdateLastModified updates the Last-Modified timestamp for a URL if it's newer.
func (t *Cassandra) UpdateLastModified(url string, newTime time.Time) error {
	if newTime.IsZero() {
		return nil // No update if no valid timestamp is provided
	}

	var storedTime time.Time
	err := t.session.Query(`SELECT last_modified_time FROM last_modified WHERE url = ?`, url).Scan(&storedTime)

	if err != nil && err != gocql.ErrNotFound {
		return err
	}

	// If the URL is new or the new timestamp is more recent, update it
	if err == gocql.ErrNotFound || newTime.After(storedTime) {
		return t.session.Query(`INSERT INTO last_modified (url, last_modified_time) VALUES (?, ?)`, url, newTime).Exec()
	}

	return nil
}

// GetLastModified fetches the last modified timestamp of a given URL
func (t *Cassandra) GetLastModified(url string) (time.Time, error) {
	var storedTime time.Time
	err := t.session.Query(`SELECT last_modified_time FROM last_modified WHERE url = ?`, url).Scan(&storedTime)

	if err != nil {
		return time.Time{}, err
	}
	return storedTime, nil
}
