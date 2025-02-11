CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR NOT NULL,
                       password VARCHAR NOT NULL,
                       balance INT NOT NULL DEFAULT 1000
);

CREATE TABLE merch (
                       id SERIAL PRIMARY KEY,
                       type VARCHAR NOT NULL,
                       price INT NOT NULL
);

CREATE TABLE purchases (
                           id SERIAL PRIMARY KEY,
                           user_id INT NOT NULL REFERENCES users(id),
                           merch_id INT NOT NULL REFERENCES merch(id),
                           amount INT NOT NULL
);

CREATE TABLE operations (
                            id SERIAL PRIMARY KEY,
                            from_user INT NOT NULL REFERENCES users(id),
                            to_user INT NOT NULL REFERENCES users(id),
                            amount INT NOT NULL
);

INSERT INTO merch (type, price) VALUES
                                    ('t-shirt', 80),
                                    ('cup', 20),
                                    ('book', 50),
                                    ('pen', 10),
                                    ('powerbank', 200),
                                    ('hoody', 300),
                                    ('umbrella', 200),
                                    ('socks', 10),
                                    ('wallet', 50),
                                    ('pink-hoody', 500);