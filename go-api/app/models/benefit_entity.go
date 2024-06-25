package models

import (
    "gorm.io/gorm"
)

type BenefitEntity struct {
    gorm.Model
    EffectiveUnitPrice     float64   `json:"effectiveUnitPrice"`
    PCToBCExchangeRate     int       `json:"pcToBcExchangeRate"`
    PCToBCExchangeRateDate string `json:"pcToBcExchangeRateDate"`
    EntitlementId          string    `json:"entitlementId"`
    EntitlementDescription string    `json:"entitlementDescription"`
    F6                     string    `json:"f6"`
    CreditPercentage       int       `json:"creditPercentage"`
    CreditType             string    `json:"creditType"`
    BenefitOrderId         string    `json:"benefitOrderId"`
    BenefitId              string    `json:"benefitId"`
    BenefitType            string    `json:"benefitType"`
    CustomerId             string    `json:"customerId"`
    CustomerName           string    `json:"customerName"`
    CustomerDomainName     string    `json:"customerDomainName"`
    CustomerCountry        string    `json:"customerCountry"`
}
