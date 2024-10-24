package numbers

import (
	"errors"
	"math"

	"go.uber.org/zap"
)

type NumbersRepository interface {
	GetIndex(value int) (int, error)
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

func (i IndexService) GetIndex(value int) (int, error) {
	idx, err := i.sortedNumbers.GetIndex(value)
	if err != nil {
		i.logger.Error("failed to get index for value",
			zap.Int("value", value),
			zap.String("error_message", err.Error()),
		)
	}

	return idx, err
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

func (s *NumbersSliceRepository) GetIndex(value int) (int, error) {
	for idx, number := range s.numbers {
		if value == number {
			s.logger.Debug("found index for value's exact match", zap.Int("value", value), zap.Int("index", idx))
			return idx, nil
		}

		numberDiff := math.Abs(float64(number - value))
		numberFoundInConformationLevel := numberDiff < float64(number)/float64(s.conformationLevel)
		if numberFoundInConformationLevel {
			s.logger.Debug("found index in conformation level", zap.Int("value", value), zap.Int("index", idx))
			return idx, nil
		}
	}

	return -1, ErrNotFound
}
