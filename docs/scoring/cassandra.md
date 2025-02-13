## cassandra.go Docs

### Connect()

Connect() is a method that helps to connect to the Cassandra DB

- Parameters:
  - hosts: []string (List of IP Addresses which have Cassandra DB Hosted)
  - keyspaceName: string (Name of the KeySpace)
- Returns:
  - Returns error (if occurred)

### Disconnect()

Disconnect() is a method that disconnect the Client from Cassandra DB

- Parameters:
- Returns:

### InitTables()

InitTables() is a method that Initialize all the Tables (Creating Columns), This is used only when the Keyspace doesn't contains all the required tables.

- Parameters:
- Returns:
  - Returns error (if occurred)
