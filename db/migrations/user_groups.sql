CREATE TABLE user_groups (
  id bigserial primary key,
  name varchar(255) not null
);

CREATE TABLE user_group_members (
  id bigserial primary key,
  user_id bigint not null references users(id),
  group_id bigint not null references user_group(id)
);

