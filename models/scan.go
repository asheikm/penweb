package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScanRequest struct {
	URL string `json:"url"`
}

type ScanResult struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ScanID    string             `bson:"scan_id"`
	ScanURL   string             `bson:"scan_url"`
	ScanDate  time.Time          `bson:"scan_date"`
	ScanData  interface{}        `bson:"scan_data"`
	CreatedAt time.Time          `bson:"created_at"`
}
