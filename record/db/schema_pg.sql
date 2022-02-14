-- Records
-- -------------------------

CREATE TABLE IF NOT EXISTS records (
    id                  SERIAL PRIMARY KEY NOT NULL,
    client              TEXT  NOT NULL, 
    total_interest      FLOAT NOT NULL,
    periodic_payment    FLOAT NOT NULL,       
    total_payment       FLOAT NOT NULL
);