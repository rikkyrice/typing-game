DROP TABLE IF EXISTS WordLists;

CREATE TABLE WordLists (
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY (START WITH 0 INCREMENT BY 1), 
    user_id VARCHAR(20) NOT NULL,
    word_list_title VARCHAR(60) NOT NULL,
    explanation VARCHAR(300), 
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (
        id
    ),

    FOREIGN KEY (user_id) REFERENCES Users(user_id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);