package numbers

import (
	"errors"
	"math"

	"go.uber.org/zap"
)

type NumbersRepository interface {
	GetIndex(value int) (Number, error)
}

type IndexService struct {
	sortedNumbers NumbersRepository
	logger        *zap.Logger
}

func NewIndexService(repo NumbersRepository, logger *zap.Logger) *IndexService {
	return &IndexService{
		sortedNumbers: repo,
		logger:        logger,
	}
}

func (i IndexService) GetIndex(value int) (Number, error) {
	number, err := i.sortedNumbers.GetIndex(value)
	if err != nil {
		i.logger.Error("failed to get index for value",
			zap.Int("value", value),
			zap.String("error_message", err.Error()),
		)
	}

	return number, err
}

var ErrNotFound = errors.New("index for given value not found")

type NumbersSliceRepository struct {
	numbers           []int
	conformationLevel int
	logger            *zap.Logger
}

func NewNumbersSliceRepository(numbers []int, conformationLevel int, logger *zap.Logger) *NumbersSliceRepository {
	return &NumbersSliceRepository{
		numbers:           numbers,
		conformationLevel: conformationLevel,
		logger:            logger,
	}
}

func (s *NumbersSliceRepository) GetIndex(value int) (Number, error) {
	leftIdx := -1
	rightIdx := len(s.numbers)

	for rightIdx > leftIdx+1 {
		middleIdx := (leftIdx + rightIdx) / 2
		if s.numbers[middleIdx] < value {
			leftIdx = middleIdx
		} else {
			rightIdx = middleIdx
		}
	}

	// TODO: add more test cases to see if these expressions could be reduced or optimized for better match in conformation level
	if rightIdx < len(s.numbers) && s.numbers[rightIdx] == value {
		return Number{Index: rightIdx, Value: value}, nil
	} else if leftIdx > -1 && s.numbers[leftIdx] == value {
		return Number{Index: leftIdx, Value: value}, nil
	} else if leftIdx > -1 && math.Abs(float64(s.numbers[leftIdx]-value)) < float64(s.numbers[leftIdx])/float64(s.conformationLevel) {
		return Number{Index: leftIdx, Value: s.numbers[leftIdx]}, nil
	} else if rightIdx < len(s.numbers) && math.Abs(float64(s.numbers[rightIdx]-value)) < float64(s.numbers[rightIdx])/float64(s.conformationLevel) {
		return Number{Index: rightIdx, Value: s.numbers[rightIdx]}, nil
	}
	return Number{Index: -1, Value: value}, ErrNotFound
}
