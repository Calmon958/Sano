# Sano Healthcare system 

## Description
This is a web-based platform that aims to bridge the gap between patients and doctors thereby ensuring easy access to quality health care. The web-app promotes Universal Health Coverage (UHC) by allowing marginalized groups of people to conveniently get health services. A patient can request for a doctor and they will get assigned the nearest available doctor to attend to them at the comfort of their home.
The app also creates an opportunity for unemployed doctors to earn some income.

## Testing
1. clone the repo

```bash
$ git clone https://github.com/Calmon958/Sano.git
```
2. Navigate to the directory

```bash
$ cd Sano
```
3. Execute the program

```bash
$ go run main.go
```
Copy the link displayed on the terminal or ctrl + click to display the platform on your web browser. Afterwards you can sign up and explore the platform.

## Features
* Authentication - there are login and signup forms for both parties.

* Profiles - there are 2 types of profiles i.e doctor's and patient's profile. Both the patient and the doctor can create accounts in the system.

* Requests - the patient can request for a doctor by creating an appointment. The doctor can also request for a followup or checkin.

* Reports - after attending to the patient the doctor has to give a report on the platform. The patient also gives their feedback.

* Database (work in progress) - all the data from the forms are stored in SQLite database.

* Payments (Work in progress) - payments are handled through the platform.

## File System
* assets - contains the templates used to capture and display data.
* database - contains the database, schematics, and initialization functions.
* web - contains the handlers and helper functions used for routing and insertion into database.
* Makefile - contains the routine tasks e.g formatting files and executing the program.

## Contributors
* Wilfred Njuguna
* Clifford Otieno
* James Muchiri
* Moffat Mokwa
* Rabin Otieno
