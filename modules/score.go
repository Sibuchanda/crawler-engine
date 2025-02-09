package modules

type Points int16

const (
	ContentImportance Points = 50
	UpdateFrequently  Points = 30
	BackLinks         Points = 20
	PageDepth         Points = 10
)

type URLScore struct {
	Score         uint64
	negativeScore uint64
	Site          string
}

// AddScore Function to Add Score into URLScore Object
func (t *URLScore) AddScore(typeof Points) {
	switch typeof {
	case ContentImportance:
		t.Score += uint64(typeof)
	case UpdateFrequently:
		t.Score += uint64(typeof)
	case BackLinks:
		t.Score += uint64(typeof)
	default:
	}
}

// SubScore Function to Subtract Score from URLScore Object
func (t *URLScore) SubScore(typeof Points) {
	if t.Score < (t.negativeScore + uint64(typeof)) {
		t.negativeScore = (t.negativeScore + uint64(typeof)) - t.Score
		t.Score = 0
		return
	}
	t.negativeScore = 0
	t.Score = t.Score - t.negativeScore

	switch typeof {
	case PageDepth:
		t.Score -= uint64(PageDepth)
	default:
	}
}
