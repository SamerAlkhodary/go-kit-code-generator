
CREATE TABLE events
(
title varchar(100) not null ,
id int unsigned auto_increment primary key ,
description text null ,
startDate datetime not null ,
endDate datetime not null ,
image varchar(200) null ,
url varchar(200) null ,
eventSort varchar(10) not null ,
timestamp   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);
CREATE TABLE locations
(
name varchar(50)  null ,
description text null ,
url varchar(100) null ,
telephone varchar(100) null ,
id --TODO: write type ,
eventId --TODO write type  NOT NULL  ,
timestamp   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);
CREATE TABLE organizers
(
name varchar(50)  null ,
logo varchar(200)  null ,
url varchar(200)  null ,
email varchar(10)  null ,
telephone varchar(20)  null ,
eventId --TODO write type  NOT NULL  ,
timestamp   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);
CREATE TABLE geos
(
latitude float null ,
longitude float null ,
location --TODO write type  NOT NULL  ,
timestamp   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);
CREATE TABLE addresss
(
streetAddress varchar(50) null ,
addressLocality varchar(20)null ,
addressCounty varchar(20)null ,
podtalCode varchar(20)null ,
location --TODO write type  NOT NULL  ,
timestamp   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);
CREATE TABLE offers
(
price varchar(20)null ,
url varchar(200)null ,
eventId --TODO write type  NOT NULL  ,
timestamp   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);
CREATE TABLE images
(
eventid --TODO write type  NOT NULL  ,
element varchar(200) null   NULL  ,
timestamp   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);