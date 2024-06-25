CREATE TABLE subscriptions (
  SubscriptionId varchar(255) DEFAULT NULL,
  SubscriptionDescription varchar(255) DEFAULT NULL,
  MpnId int(11) DEFAULT NULL,
  Tier2MpnId int(11) DEFAULT NULL,
  InvoiceNumber varchar(255) DEFAULT NULL,
  PartnerId varchar(255) DEFAULT NULL
);