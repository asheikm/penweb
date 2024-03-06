package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"penweb/models"

	"go.mongodb.org/mongo-driver/bson"

	//"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScanService interface {
	PerformScan(*models.ScanRequest) error
	GetCompletedScanResult(string) (*models.ScanResult, error)
}

type dbScanService struct {
	db *mongo.Database
}

func NewDbScanService(db *mongo.Database) ScanService {
	return &dbScanService{db: db}
}

func (r *dbScanService) PerformScan(s *models.ScanRequest) error {
	payload, err := json.Marshal(s)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://owasp-api-service/start-scan", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to start scan")
	}

	return nil
}

func (r *dbScanService) GetCompletedScanResult(scanID string) (*models.ScanResult, error) {
	collection := r.db.Collection("scan_results")

	filter := bson.M{"scan_id": scanID}
	var result models.ScanResult
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
