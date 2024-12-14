CREATE TABLE patients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    phone_number TEXT NOT NULL UNIQUE,
    id_number TEXT NOT NULL,
    otp_code TEXT,
    otp_expiry DATETIME
);

-- create patient registration table
CREATE TABLE patient_registration (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    Age INTEGER NOT NULL,
    id_number TEXT NOT NULL UNIQUE,
    physical_address TEXT NOT NULL,
    phone_number TEXT NOT NULL CHECK(LENGTH(phone_number) = 10 AND phone_number GLOB '[0-9]*')
);


CREATE TABLE appointments (
    appointment_id INTEGER PRIMARY KEY AUTOINCREMENT,
    patient_id TEXT NOT NULL UNIQUE,
    doctor_id TEXT NOT NULL UNIQUE,
    doctor_name TEXT NOT NULL,
    appointment_details TEXT NOT NULL,
    appointment_date TEXT NOT NULL,
    appointment_time TEXT NOT NULL,
    document TEXT, --optional field for document upload
    FOREIGN KEY( patient_ID) REFERENCES patient_registration (id_number),
    FOREIGN KEY (doctor_ID) REFERENCES doctor_registration (license)


)

