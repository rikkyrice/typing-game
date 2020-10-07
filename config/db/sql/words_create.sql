DROP TABLE IF EXISTS Words;

CREATE TABLE Words (
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY (START WITH 0 INCREMENT BY 1), 
    word_list_id INT NOT NULL,
    word VARCHAR(60) NOT NULL, 
    meaning VARCHAR(100),
    explanation VARCHAR(300),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (
        id
    ),

    FOREIGN KEY (word_list_id) REFERENCES WordLists(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);