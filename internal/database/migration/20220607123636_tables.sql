-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS userservice.otp (
externalID uuid  PRIMARY KEY DEFAULT uuid_generate_v4(),
otp_keys text [10]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS otp CASCADE;

-- +goose StatementEnd
