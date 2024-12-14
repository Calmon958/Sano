CREATE TABLE users {
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(10) NOT NULL UNIQUE,
    ID_Number VARCHAR(8) NOT NULL,
    otp_code VARCHAR(6),
    otp_expiry TIMESTAMP
};