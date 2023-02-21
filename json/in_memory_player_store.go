package json

// init empty player store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// collect data player
type InMemoryPlayerStore struct {
	store map[string]int
}

// record player win
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// retreivescores for player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
