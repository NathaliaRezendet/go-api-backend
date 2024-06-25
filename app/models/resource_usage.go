package models

type ResourceUsage struct {
    UsageDate        string `json:"usageDate"`
    MeterType        string `json:"meterType"`
    MeterCategory    string `json:"meterCategory"`
    MeterId          string `json:"meterId"`
    MeterSubCategory string `json:"meterSubCategory"`
    MeterName        string `json:"meterName"`
    MeterRegion      string `json:"meterRegion"`
    Unit             string `json:"unit"`
    ResourceLocation string `json:"resourceLocation"`
    ConsumedService  string `json:"consumedService"`
    ResourceGroup    string `json:"resourceGroup"`
    ResourceURI      string `json:"resourceURI"`
}
