package store

// Migrator ...
type Migrator interface {
	MigrateDB() error
	ReverseDB() error
}
