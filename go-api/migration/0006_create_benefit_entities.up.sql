CREATE TABLE benefit_entities (
  EffectiveUnitPrice double DEFAULT NULL,
  PCToBCExchangeRate int(11) DEFAULT NULL,
  PCToBCExchangeRateDate time DEFAULT NULL,
  EntitlementId varchar(255) DEFAULT NULL,
  EntitlementDescription varchar(255) DEFAULT NULL,
  f6 varchar(255) DEFAULT NULL,
  CreditPercentage int(11) DEFAULT NULL,
  CreditType varchar(255) DEFAULT NULL,
  BenefitOrderId varchar(255) DEFAULT NULL,
  BenefitId varchar(255) DEFAULT NULL,
  BenefitType varchar(255) DEFAULT NULL,
  CustomerId varchar(255) DEFAULT NULL,
  CustomerName varchar(255) DEFAULT NULL,
  CustomerDomainName varchar(255) DEFAULT NULL,
  CustomerCountry varchar(255) DEFAULT NULL
);