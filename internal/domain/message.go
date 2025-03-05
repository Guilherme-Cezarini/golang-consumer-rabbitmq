package domain

import "time"

type Message struct {
	Message    string    `bson:"message" json:"message"`
	CompanyID  string    `bson:"company_id" json:"company_id"`
	CampaignID string    `bson:"campaign_id" json:"campaign_id"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
}
