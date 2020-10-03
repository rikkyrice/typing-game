DROP TABLE IF EXISTS WordLists;

CREATE TABLE WordLists (
    id INT AUTO_INCREMENT PRIMARY KEY, 
    user_id VARCHAR(20) NOT NULL,
    word_list_title VARCHAR(60) NOT NULL,
    explanation VARCHAR(300), 
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    FOREIGN KEY (user_id) REFERENCES Users(user_id)
        ON UPDATE NO ACTION
        ON DELETE cascade
);