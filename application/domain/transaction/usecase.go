package transaction

import (
	"context"
	"log"
	"simple-go/pkg/response"
)

type Repository interface {
	TransactionFreqCheck(ctx context.Context) ([]FreqCheckResponse, error)
}

type service struct {
	repository Repository
}

func (s service) getDataTransactionFreq(ctx context.Context, resp chan<- []FreqCheckResponse) {
	var data []FreqCheckResponse
	data, _ = s.repository.TransactionFreqCheck(ctx)

	for _, val := range data {
		if val.Total > 8 {
			val.Score = 90
		} else if val.Total >= 7 || val.Total <= 8 {
			val.Score = 80
		} else if val.Total == 6 {
			val.Score = 70
		} else if val.Total == 5 {
			val.Score = 50
		} else if val.Total < 5 {
			val.Score = 40
		}
	}

	resp <- data
}
func (s service) getDataAmountCheck(ctx context.Context, resp chan<- []FreqCheckResponse) {
	var data []FreqCheckResponse
	resp <- data
}
func (s service) getDataPatterCheck(ctx context.Context, resp chan<- []FreqCheckResponse) {
	var data []FreqCheckResponse
	resp <- data
}

func (s service) GetFraudDetection(ctx context.Context) ([]ResultDetectionData, response.ErrorResponse) {
	//get transaction freq
	transactionFreqCh := make(chan []FreqCheckResponse)
	go s.getDataTransactionFreq(ctx, transactionFreqCh)
	//get amount check
	amountCheckCh := make(chan []FreqCheckResponse)
	go s.getDataAmountCheck(ctx, amountCheckCh)
	//get pattern check
	patternCheckCh := make(chan []FreqCheckResponse)
	go s.getDataPatterCheck(ctx, patternCheckCh)

	//map to array with user id
	transactionFreq := <-transactionFreqCh
	log.Println(transactionFreq)

	return []ResultDetectionData{}, *response.NotError()
}

func NewService(repo Repository) service {
	return service{
		repository: repo,
	}
}
