CREATE TYPE owner_type AS ENUM (
  'user_group',
  'user'
);

CREATE TABLE task_groups (
  id bigserial primary key,
  name varchar(255) not null,
  owner_type owner_type not null,
  owner_user_group_id bigint references user_group(id),
  owner_user_id bigint references users(id)
);