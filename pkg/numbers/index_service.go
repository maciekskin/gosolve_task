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
	for idx, number := range s.numbers {
		if value == number {
			s.logger.Debug("found index for value's exact match", zap.Int("value", value), zap.Int("index", idx))
			return Number{Index: idx, Value: number}, nil
		}

		numberDiff := math.Abs(float64(number - value))
		numberFoundInConformationLevel := numberDiff < float64(number)/float64(s.conformationLevel)
		if numberFoundInConformationLevel {
			// TODO:
			// currently there is a bug that for big numbers we're not finding exact match because we find a value that fits the conformation level
			// use binary search to have better lookup performance and to return correct indexes!
			s.logger.Debug("found index in conformation level", zap.Int("value", value), zap.Int("index", idx))
			return Number{Index: idx, Value: number}, nil
		}
	}

	return Number{Index: -1, Value: value}, ErrNotFound
}
