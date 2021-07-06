CREATE TABLE cities
(
    id           SERIAL PRIMARY KEY,
    country_name VARCHAR(25),
    city_name    VARCHAR(25)
);

CREATE TABLE flats
(
    id           SERIAL PRIMARY KEY,
    street       VARCHAR(50) NOT NULL,
    house_number VARCHAR(5)  NOT NULL,
    room_number  INT         NOT NULL,
    description TEXT,
    city_id      INT         NOT NULL REFERENCES cities (id)
);

INSERT INTO cities (country_name, city_name) VALUES ('Belarus', 'Minsk');