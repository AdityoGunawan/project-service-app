create database service_app;

use service_app;

create table users (
no_rekening varchar(100) primary key not null,
nomor_telepon varchar(100) not null,
nama varchar(100),
password varchar(100)not null,
saldo varchar(100)default null,
gender enum ("laki-laki, perempuan"),
addres varchar(200)default null
);

create table topup (
topup_id varchar(100)primary key not null,
no_rekening varchar(100) not null,
nominal_topup varchar(100) not null,
history_topup datetime default current_timestamp
);
ALTER TABLE service_app.topup DROP FOREIGN KEY FK_TOPUP_USERS;
ALTER TABLE service_app.topup ADD CONSTRAINT FK_TOPUP_USERS FOREIGN KEY (no_rekening) REFERENCES service_app.users(no_rekening) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE service_app.topup DROP FOREIGN KEY topup_FK_1;
ALTER TABLE service_app.topup DROP FOREIGN KEY topup_FK;
ALTER TABLE service_app.topup DROP FOREIGN KEY FK_TOPUP_USERS;
ALTER TABLE service_app.topup ADD CONSTRAINT FK_TOPUP_USERS FOREIGN KEY (no_rekening) REFERENCES service_app.users(no_rekening) ON DELETE CASCADE ON UPDATE CASCADE;



create table transfer (
transfer_id varchar(100)primary key not null,
no_rekening_pengirim varchar(100) not null,
no_rekening_penerima varchar(100) not null,
nominal_transfer varchar(100) not null,
history_transfer datetime default current_timestamp,
CONSTRAINT FK_TRANSFER_USERS FOREIGN KEY (no_rekening_pengirim) REFERENCES service_app.users(no_rekening) ON DELETE CASCADE ON UPDATE cascade,
CONSTRAINT FK_TRANSFERS_USERS FOREIGN KEY (no_rekening_penerima) REFERENCES service_app.users(no_rekening) ON DELETE CASCADE ON UPDATE CASCADE
);