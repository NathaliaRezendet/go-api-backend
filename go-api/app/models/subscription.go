package models

type Subscription struct {
    SubscriptionId          string `json:"subscriptionId"`
    SubscriptionDescription string `json:"subscriptionDescription"`
    MpnId                   int    `json:"mpnId"`
    Tier2MpnId              int    `json:"tier2MpnId"`
    InvoiceNumber           string `json:"invoiceNumber"`
    PartnerId               string `json:"partnerId"`
}