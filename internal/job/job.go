package job

// Job is the planned API payload and persisted job record.
// TODO: choose ID generation, validation limits, and retention policy.
type Job struct {
    ID string `json:"id"`
    Text string `json:"text"`
    Status string `json:"status"`
    WordCount *int `json:"word_count,omitempty"`
}
