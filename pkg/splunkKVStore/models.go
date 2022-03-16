package kvstore

// RecentRecord is saved in splunk's mongodb
// it is in relation to the current snake
type RecentRecord struct {
	GameID       string
	SnakeSAIID   string
	SnakeGSID    string
	Turns        int
	ScoreChange  string
	PointChange  string
	TierChange   string
	Result       string
	SnakeTimeout int
}

// GameOutcomeRecord is based off the 'engine' endpoint and
// holds the result for each snake (from 'LastFrame')
type GameOutcomeRecord struct {
	GameID    string
	SnakeGSID string
	Name      string
	Death     struct {
		Cause        string
		Turn         int
		EliminatedBy string
	}
}
