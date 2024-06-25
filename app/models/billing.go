package models

type Billing struct {
    ID                 uint    `json:"id" gorm:"primaryKey"`
    ChargeType         string  `json:"chargeType"`
    UnitPrice          float64 `json:"unitPrice"`
    Quantity           float64 `json:"quantity"`
    UnitType           string  `json:"unitType"`
    BillingPreTaxTotal float64 `json:"billingPreTaxTotal"`
    BillingCurrency    string  `json:"billingCurrency"`
    PricingPreTaxTotal float64 `json:"pricingPreTaxTotal"`
    PricingCurrency    string  `json:"pricingCurrency"`
}
