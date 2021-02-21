DROP TABLE IF EXISTS Words;

CREATE TABLE Words (
    id CHAR(36) NOT NULL, 
    word_list_id CHAR(36) NOT NULL,
    word VARCHAR(60) NOT NULL,
    yomi VARCHAR(120),
    meaning VARCHAR(300),
    m_yomi VARCHAR(600),
    explanation VARCHAR(1024),
    is_remembered BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (
        id
    ),

    FOREIGN KEY (word_list_id) REFERENCES WordLists(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);