CREATE TABLE resources_usage (
  UsageDate varchar(255) DEFAULT NULL,
  MeterType varchar(255) DEFAULT NULL,
  MeterCategory varchar(255) DEFAULT NULL,
  MeterId varchar(255) DEFAULT NULL,
  MeterSubCategory varchar(255) DEFAULT NULL,
  MeterName varchar(255) DEFAULT NULL,
  MeterRegion varchar(255) DEFAULT NULL,
  Unit varchar(255) DEFAULT NULL,
  ResourceLocation varchar(255) DEFAULT NULL,
  ConsumedService varchar(255) DEFAULT NULL,
  ResourceGroup varchar(255) DEFAULT NULL,
  ResourceURI varchar(255) DEFAULT NULL
);