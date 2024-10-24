package numbers

import (
	"errors"
	"math"
)

// TODO:
// - implement service with single method to look for index of given value with 10% level of conformation
// - add unit tests
// - move service/repository to separate package

type NumbersRepository interface {
	GetIndex(value int) (int, error)
}

type IndexService struct {
	sortedNumbers NumbersRepository
	// TODO: add logger
}

func NewIndexService(repo NumbersRepository) *IndexService {
	return &IndexService{
		sortedNumbers: repo,
	}
}

func (i IndexService) GetIndex(value int) (int, error) {
	idx, err := i.sortedNumbers.GetIndex(value)
	if err != nil {
		// TODO: log error at service level
	}

	return idx, err
}

var ErrNotFound = errors.New("index for given value not found")

type NumbersSliceRepository struct {
	numbers           []int
	conformationLevel int
}

func NewNumbersSliceRepository(numbers []int, conformationLevel int) *NumbersSliceRepository {
	return &NumbersSliceRepository{
		numbers:           numbers,
		conformationLevel: conformationLevel,
	}
}

func (s *NumbersSliceRepository) GetIndex(value int) (int, error) {
	// -1 means that index wasn't found for either exact value or in conformation level
	closestNumberIdx := -1
	closestNumberDiff := float64(value)

	for idx, number := range s.numbers {
		if value == number {
			closestNumberIdx = idx
			break
		}

		numberDiff := math.Abs(float64(number - value))
		numberFoundInConformationLevel := numberDiff < float64(number)/float64(s.conformationLevel)
		if numberFoundInConformationLevel && numberDiff < closestNumberDiff {
			// Note - to be decided during code review:
			// for better efficiency we could omit finding closestNumberIdx
			// and simply return the first match below conformation level
			closestNumberIdx = idx
			closestNumberDiff = numberDiff
		}
	}

	if closestNumberIdx < 0 {
		return closestNumberIdx, ErrNotFound
	}
	return closestNumberIdx, nil
}
