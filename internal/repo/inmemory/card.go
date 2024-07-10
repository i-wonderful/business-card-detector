package inmemory

import (
	"card_detector/internal/model"
	"sort"
	"sync"
)

type CardRepo struct {
	store  map[int64]model.Card
	lastId int64
	mu     sync.RWMutex
}

func NewCardRepo() *CardRepo {
	return &CardRepo{
		store: make(map[int64]model.Card),
		mu:    sync.RWMutex{},
	}
}

func (r *CardRepo) GetPerson(id int64) (model.Card, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.store[id], nil
}

func (r *CardRepo) Save(card model.Card) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastId++
	card.Id = r.lastId
	r.store[r.lastId] = card

	return nil
}

func (r *CardRepo) GetAll() []model.Card {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var cards []model.Card
	for _, c := range r.store {
		cards = append(cards, c)
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[j].UploadedAt.Before(cards[i].UploadedAt)
	})

	return cards
}
