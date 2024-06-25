package models

type Client struct {
    ID                 uint   `json:"id" gorm:"primaryKey"`
    CustomerId         string `json:"customerId"`
    CustomerName       string `json:"customerName"`
    CustomerDomainName string `json:"customerDomainName"`
    CustomerCountry    string `json:"customerCountry"`
}
