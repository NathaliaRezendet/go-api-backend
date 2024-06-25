CREATE TABLE billings (
  ChargeType varchar(255) DEFAULT NULL,
  UnitPrice double DEFAULT NULL,
  Quantity double DEFAULT NULL,
  UnitType varchar(255) DEFAULT NULL,
  BillingPreTaxTotal double DEFAULT NULL,
  BillingCurrency varchar(255) DEFAULT NULL,
  PricingPreTaxTotal double DEFAULT NULL,
  PricingCurrency varchar(255) DEFAULT NULL
);