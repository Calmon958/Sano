CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    phone_number TEXT NOT NULL UNIQUE,
    ID_Number TEXT NOT NULL,
    otp_code TEXT,
    otp_expiry DATETIME
);