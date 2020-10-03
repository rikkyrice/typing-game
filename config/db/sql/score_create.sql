DROP TABLE IF EXISTS Score;

CREATE TABLE Score (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    WordList INT NOT NULL,
    PlayCount INT NOT NULL,
    ClearTime TIME NOT NULL,
    ClearTypeCount INT NOT NULL,
    MissTypeCount INT NOT NULL,
    CorrectRate DOUBLE NOT NULL,
    LastPlay_at DATETIME NOT NULL
);