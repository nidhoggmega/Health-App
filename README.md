# Healthcare App

## Design Considerations

This solution is designed to improve access to healthcare services by providing two microservices: "appointment_service" and "patient_records_service". The microservices are built using Go language, follow microservices architecture, and utilize MySQL for data storage. Key design considerations include:

- **Scalability:** The microservices are designed to scale independently, allowing easy horizontal scaling based on demand.
- **Containerization:** Each microservice is containerized using Docker, ensuring consistent deployment across different environments.
- **Error Handling:** Error handling has been implemented in each function to provide meaningful error messages and responses.
- **RESTful API:** The microservices expose RESTful APIs for creating, retrieving, and deleting appointments and patient records.
- **Database Connectivity:** The microservices connect to a MySQL database to store and retrieve data.

## Architecture Diagram

![Unnamed File](https://github.com/nidhoggmega/Health-App/assets/92702429/ccd189a0-cd93-4dc8-9d81-20f39437b345)



## Instructions for Setup and Running

### 1. Set Up MySQL Database

- Open MySQL Workbench and create a new database named "healthcare_db".
- Execute the SQL script in the provided code to create tables for "appointments" and "patient_records".

- The SQL setup code is given in the master branch.

### 2. Run Microservices

#### For "appointment_service" Microservice:

##### Running:
cd path/to/appointment_service
go run main.go

#### For "patient_records_service" Microservice:

##### Running:
cd path/to/patient_records_service
go run main.go

### 3. Test Microservices

#### Testing appointment_service Microservice:

##### Create an appointment:

curl -X POST -H "Content-Type: application/json" -d '{"appointment": "Dentist"}' http://localhost:8080/appointments

##### Expected Output:

"Appointment created successfully"

##### Retrieve all appointments:

curl http://localhost:8080/appointments

##### Expected Output:

"[{"id":1,"appointment":"Dentist"}]"

##### Delete an appointment (replace {id} with the actual appointment ID):

curl -X DELETE http://localhost:8080/appointments/{id}

##### Expected Output:

"Appointment deleted successfully"

#### Testing patient_records_service Microservice:

##### Create a patient record:

curl -X POST -H "Content-Type: application/json" -d '{"patient_record": "Blood Pressure: 120/80"}' http://localhost:8081/patient-records

##### Expected Output:

"Patient record created successfully"

##### Retrieve all patient records:

curl http://localhost:8081/patient-records

##### Expected Output:

"[{"id":1,"patient_record":"Blood Pressure: 120/80"}]"

##### Delete a patient record (replace {id} with the actual record ID):

curl -X DELETE http://localhost:8081/patient-records/{id}

"Patient record deleted successfully"

### 4. Build & Run the Docker Containers

##### Build the Docker images for both services:

docker build -t appointment-service ./appointment_service
docker build -t patient-records-service ./patient_records_service

##### Run the Docker containers:

docker run -p 8080:8080 appointment-service
docker run -p 8081:8081 patient-records-service

## Note:

##### Replace placeholders such as path/to/appointment_service with actual paths when running on different devices.
##### Make sure to replace {id} in DELETE requests with actual IDs retrieved from GET requests.
