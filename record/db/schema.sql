-- Records
-- -------------------------

CREATE TABLE IF NOT EXISTS records (
    id                  INTEGER PRIMARY KEY autoincrement,
    client              TEXT  NOT NULL, 
    total_interest      FLOAT NOT NULL,
    periodic_payment    FLOAT NOT NULL,       
    total_payment       FLOAT NOT NULL
);