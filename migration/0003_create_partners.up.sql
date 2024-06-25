CREATE TABLE partners (
  id int(11) NOT NULL AUTO_INCREMENT,
  PartnerId varchar(255) DEFAULT NULL,
  PartnerName varchar(255) DEFAULT NULL,
  PublisherName varchar(255) DEFAULT NULL,
  PublisherId varchar(255) DEFAULT NULL,
  PartnerEarnedCreditPercentage int(11) DEFAULT NULL,
  PRIMARY KEY (id)
);