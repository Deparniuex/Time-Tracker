CREATE TABLE users (
    id bigserial PRIMARY KEY,
    surname text NOT NULL,
    first_name text NOT NULL,
    patronymic text NOT NULL,
    address text NOT NULL,
    passport_serie int NOT NUll,
    passport_number int NOT NULL
);

INSERT INTO users (surname, first_name, patronymic, address, passport_serie, passport_number) 
VALUES ('Иванов', 'Иван', 'Иванович', 'г. Москва, ул. Ленина, д. 5, кв. 1', 1234, 567890);