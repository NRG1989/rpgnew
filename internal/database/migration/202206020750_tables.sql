-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS userservice.client_o2o_otp (
externalID uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
mobile_phone varchar(13) REFERENCES userservice.client(mobile_phone)
);

CREATE TABLE IF NOT EXISTS userservice.otp (
externalID uuid REFERENCES userservice.client_o2o_otp(externalID),
otp_keys integer [10]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS userservice.otp   CASCADE;
DROP TABLE IF EXISTS userservice.client_o2o_otp;
DROP EXTENSION IF EXISTS "uuid-ossp";

-- +goose StatementEnd
