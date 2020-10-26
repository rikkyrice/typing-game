DROP TABLE IF EXISTS Token;

CREATE TABLE Token (
    token VARCHAR(500) NOT NULL, 
    user_id VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    PRIMARY KEY (
        token
    ),

    FOREIGN KEY (user_id) REFERENCES Users(user_id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);