package db

const patients_table  string = `
CREATE TABLE patients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    phone_number TEXT NOT NULL UNIQUE,
    id_number TEXT NOT NULL,
    otp_code TEXT,
    otp_expiry DATETIME
); 
`

//-- create patient registration table
const patient_reg_table  string = `
CREATE TABLE patient_registration (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    Age INTEGER NOT NULL,
    id_number TEXT NOT NULL UNIQUE,
    physical_address TEXT NOT NULL,
    phone_number TEXT NOT NULL CHECK(LENGTH(phone_number) = 10 AND phone_number GLOB '[0-9]*')
);
`


const apppointments_table string = `
CREATE TABLE appointments (
    appointment_id INTEGER PRIMARY KEY AUTOINCREMENT,
    patient_id TEXT NOT NULL UNIQUE,
    doctor_id TEXT NOT NULL UNIQUE,
    doctor_name TEXT NOT NULL,
    appointment_details TEXT NOT NULL,
    appointment_date TEXT NOT NULL,
    appointment_time TEXT NOT NULL,
    document TEXT, --optional field for document upload
    FOREIGN KEY( patient_id) REFERENCES patient_registration (id_number),
    FOREIGN KEY (doctor_id) REFERENCES doctor_registration (license)


);
`

const doctor_reg_table string = `
CREATE TABLE doctor_registration (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    title TEXT NOT NULL
    doctor_id_number TEXT NOT NULL UNIQUE,
    doctor_license TEXT NOT NULL UNIQUE,
    physical_address TEXT NOT NULL,
    organization TEXT,
    phone_number TEXT NOT NULL CHECK(LENGTH(phone_number) = 10 AND phone_number GLOB '[0-9]*')
);
`

const doctor_report_table string =`
CREATE TABLE doctor_report (
    report_id INTEGER PRIMARY KEY AUTOINCREMENT,
    report_name TEXT NOT NULL,
    report_type TEXT NOT NULL,
    file_path TEXT NOT NULL,
    upload_date TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    doctor_id TEXT NOT NULL,
    FOREIGN KEY (doctor_id) REFERENCES doctor_registration (doctor_license)
);
`

var Tables = []string{patients_table, patient_reg_table, apppointments_table, doctor_reg_table, doctor_report_table}