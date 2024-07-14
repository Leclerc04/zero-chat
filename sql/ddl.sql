CREATE TABLE Contacts
(
    id         BIGINT    NOT NULL,
    owner_id   BIGINT    NOT NULL,
    contact_id BIGINT    NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE Message (
                         id BIGINT NOT NULL,
                         send_id BIGINT NOT NULL,
                         receive_id BIGINT NOT NULL,
                         msg TEXT NOT NULL,
                         created_at TIMESTAMP NOT NULL,
                         updated_at TIMESTAMP,
                         deleted_at TIMESTAMP,
                         PRIMARY KEY (id)
);

CREATE TABLE User (
                      id BIGINT NOT NULL,
                      email VARCHAR(255) NOT NULL,
                      password VARCHAR(255) NOT NULL,
                      nickname VARCHAR(255) NOT NULL,
                      created_at TIMESTAMP NOT NULL,
                      updated_at TIMESTAMP,
                      deleted_at TIMESTAMP,
                      PRIMARY KEY (id)
);
