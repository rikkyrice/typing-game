DROP TABLE IF EXISTS Scores;

CREATE TABLE Scores (
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY (START WITH 0 INCREMENT BY 1),
    word_list_id INT NOT NULL,
    play_count INT NOT NULL,
    clear_type_count INT NOT NULL,
    miss_type_count INT NOT NULL,
    correct_rate DOUBLE NOT NULL,
    played_at DATETIME NOT NULL,
    PRIMARY KEY (
        id
    ),

    FOREIGN KEY (word_list_id) REFERENCES WordLists(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);