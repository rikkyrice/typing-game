DROP TABLE IF EXISTS User;

CREATE TABLE User (
    ID INT AUTO_INCREMENT PRIMARY KEY, 
    Name VARCHAR(50) NOT NULL, 
    Mail VARCHAR(100) NOT NULL,
    Password VARCHAR(100) NOT NULL,
    Created_at DATETIME NOT NULL
);