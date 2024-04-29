CREATE TABLE IF NOT EXISTS personas (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    edad INTEGER NOT NULL,
    country_code VARCHAR(10) NOT NULL
);

CREATE DATABASE IF NOT EXISTS personas_test;


\c personas_test;

-- Crear la tabla personas si no existe
CREATE TABLE IF NOT EXISTS personas (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    edad INTEGER NOT NULL,
    country_code VARCHAR(10) NOT NULL
);