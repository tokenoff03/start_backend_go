-- +goose Up
-- +goose StatementBegin
CREATE TABLE room (  
    id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    number VARCHAR(255),
    type VARCHAR(255),
    description VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE room;
-- +goose StatementEnd
