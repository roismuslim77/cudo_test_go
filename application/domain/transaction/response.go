package transaction

type FreqCheckResponse struct {
	Total   int    `json:"total"`
	UserId  int    `json:"user_id"`
	OrderId string `json:"order_id"`
	Score   int    `json:"score"`
}

type DetectionResultData struct {
	IsSuspicious    bool     `json:"is_suspicious"`
	ConfidenceScore float64  `json:"confidence_score"`
	Triggers        []string `json:"triggers"`
}

type ResultDetectionData struct {
	TransactionId   string  `json:"transaction_id"`
	FraudScore      float64 `json:"fraud_score"`
	RiskLevel       string  `json:"risk_level"`
	DetectionResult struct {
		FrequencyCheck DetectionResultData `json:"frequency_check"`
		AmountCheck    DetectionResultData `json:"amount_check"`
		PatternCheck   DetectionResultData `json:"pattern_check"`
	} `json:"detection_result"`
}
