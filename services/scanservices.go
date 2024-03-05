package services

import (
	"penweb/models"

	//"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScanService interface {
	PerformScan(*models.ScanRequest) error
	GetCompletedScan(string) (*models.ScanResult, error)
}

type dbScanService struct {
	db *mongo.Database
}

func NewDbScanService(db *mongo.Database) ScanService {
	return &dbScanService{db: db}
}

func (r *dbScanService) PerformScan(s *models.ScanRequest) error {
	return nil

}

func (r *dbScanService) GetCompletedScan(scanId string) (*models.ScanResult, error) {
	return nil, nil

}
