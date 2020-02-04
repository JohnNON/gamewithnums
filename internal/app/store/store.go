package store

// Store - интерфейс, описывающий хранилище
type Store interface {
	User() UserRepository
}
