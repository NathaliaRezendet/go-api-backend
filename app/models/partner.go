package models

type Partner struct {
    ID                        uint   `json:"id" gorm:"primaryKey"`
    PartnerId                 string `json:"partnerId"`
    PartnerName               string `json:"partnerName"`
    PublisherName             string `json:"publisherName"`
    PublisherId               string `json:"publisherId"`
    PartnerEarnedCreditPercentage int `json:"partnerEarnedCreditPercentage"`
}
