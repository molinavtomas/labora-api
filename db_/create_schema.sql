CREATE TABLE IF NOT EXISTS db.personas (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    edad INTEGER NOT NULL,
    country_code VARCHAR(10) NOT NULL
);