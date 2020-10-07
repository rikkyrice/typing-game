DROP TABLE IF EXISTS Users;

CREATE TABLE Users (
    user_id VARCHAR(20) NOT NULL, 
    mail VARCHAR(100) NOT NULL,
    password VARCHAR(20) NOT NULL,
    created_at DATE NOT NULL,
    PRIMARY KEY (
        user_id
    )
);