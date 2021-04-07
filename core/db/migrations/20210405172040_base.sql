-- +goose Up
-- +goose StatementBegin
create table users
(
    id         integer
        primary key,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    name       text
);

create
index idx_users_deleted_at
	on users (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
