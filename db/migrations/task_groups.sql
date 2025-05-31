CREATE TABLE task_groups (
  id bigserial primary key,
  name varchar(255) not null,
  owner_user_id bigint references users(id)
);

CREATE TABLE task_group_members (
  id bigserial primary key,
  task_group_id bigint not null references task_groups(id),
  user_id bigint not null references users(id)
);