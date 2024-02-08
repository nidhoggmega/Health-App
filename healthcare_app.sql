CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'localhost';

-- For appointments table
CREATE TABLE appointments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    appointment VARCHAR(255) NOT NULL
);

-- For patient_records table
CREATE TABLE patient_records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    patient_record VARCHAR(255) NOT NULL
);
