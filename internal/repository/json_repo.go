package repository

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/Kotoninja/expense-tracker/pkg/expense"
)

type JSONRepository struct {
	mu        sync.RWMutex
	filePath  string
	data      map[int]expense.Expense
	increment int
}

func NewJSONRepository(filePath string) (*JSONRepository, error) {
	repo := &JSONRepository{
		filePath:  filePath,
		data:      make(map[int]expense.Expense),
		increment: 1,
	}

	if err := repo.load(); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *JSONRepository) load() error {
	file, err := os.Open(r.filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&r.data); err != nil {
		return err
	}

	for k := range r.data {
		if k >= r.increment {
			r.increment = k + 1
		}
	}

	return nil
}

func (r *JSONRepository) save() error {
	file, err := os.OpenFile(r.filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(r.data)
}

func (r *JSONRepository) Add(exp *expense.Expense) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.increment
	exp.ID = id
	r.data[id] = *exp
	r.increment++

	if err := r.save(); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *JSONRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.data, id)
	return r.save()
}

func (r *JSONRepository) Update(id int, exp *expense.Expense) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("expense not found")
	}

	exp.ID = id
	r.data[id] = *exp
	return r.save()
}

func (r *JSONRepository) FindByID(id int) (*expense.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	exp, ok := r.data[id]
	if !ok {
		return nil, errors.New("expense not found")
	}

	return &exp, nil
}

func (r *JSONRepository) FindAll() ([]expense.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]expense.Expense, 0, len(r.data))
	for _, exp := range r.data {
		result = append(result, exp)
	}

	return result, nil
}

func (r *JSONRepository) GetTotal(month int) int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	total := 0
	for _, exp := range r.data {
		if month == 0 || int(exp.Date.Month()) == month {
			total += exp.Amount
		}
	}
	return total
}
