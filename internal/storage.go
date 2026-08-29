package internal

import (
	"encoding/json"
	"errors"
	"os"
	"sort"
	"strconv"
	"time"

	expense "github.com/Kotoninja/expense-tracker/pkg"
)

var (
	StorageIO *Storage
)

type Storage struct {
	increment int
	data      map[int]expense.Expense
	filePath  string
}

func (s *Storage) load() error {
	file, err := os.Open(s.filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	} else if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&s.data); err != nil {
		return err
	}

	for k := range s.data {
		if k >= s.increment {
			s.increment = k + 1
		}
	}

	return nil
}

func (s *Storage) save() error {
	file, err := os.OpenFile(s.filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err = encoder.Encode(s.data); err != nil {
		return err
	}

	return nil
}

func NewStorage(filePath string) (*Storage, error) {
	newStore := &Storage{increment: 1, filePath: filePath, data: map[int]expense.Expense{}}

	if err := newStore.load(); err != nil {
		return nil, err
	}

	return newStore, nil
}

func (s *Storage) Add(description *string, amount *int, timestamp *time.Time) (*int, error) {
	newExp := expense.NewExpense(*description, *amount, timestamp)
	s.data[s.increment] = *newExp
	defer func() { s.increment++ }()
	if err := s.save(); err != nil {
		return nil, err
	}
	return &s.increment, nil
}

func (s *Storage) Delete(id int) error {
	delete(s.data, id)
	if err := s.save(); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Update(id int, description string, amount int, timstamp time.Time) error {
	obj, ok := s.data[id]
	if !ok {
		return errors.New("Object not found")
	}

	if description != "" {
		obj.Description = description
	}
	if amount != 0 {
		obj.Amount = amount
	}
	if !timstamp.IsZero() {
		obj.Date = timstamp
	}

	s.data[id] = obj
	s.save()
	return nil
}

func (s *Storage) List() [][]string {
	result := [][]string{}

	for key, value := range s.data {
		result = append(result, []string{
			strconv.Itoa(key),
			value.Description,
			strconv.Itoa(value.Amount),
			value.Date.Format(time.DateOnly),
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

func (s *Storage) Summary() int {
	var answer int

	for _, exp := range s.data {
		answer += exp.Amount
	}

	return answer
}
