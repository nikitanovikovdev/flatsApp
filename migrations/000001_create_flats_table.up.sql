CREATE TABLE cities (
    id           SERIAL PRIMARY KEY,
    country_name VARCHAR(25),
    city_name    VARCHAR(25)
);

CREATE TABLE flats (
    id           SERIAL PRIMARY KEY,
    street       VARCHAR(50) NOT NULL CHECK (street <> ' '),
    house_number VARCHAR(5)  NOT NULL CHECK (house_number <> ' '),
    room_number  INT         NOT NULL CHECK (room_number > 0),
    description TEXT,
    city_id      INT         NOT NULL REFERENCES cities (id)
);

