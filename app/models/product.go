package models


type Product struct {
    ProductId      string    `json:"productId"`
    SkuId          string    `json:"skuId"`
    AvailabilityId string    `json:"availabilityId"`
    SkuName        string    `json:"skuName"`
    ProductName    string    `json:"productName"`
    ChargeStartDate string `json:"chargeStartDate"`
    ChargeEndDate   string `json:"chargeEndDate"`
}
