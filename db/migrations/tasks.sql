CREATE TYPE task_status AS ENUM (
  'in_progress',
  'done',
  'canceled'
);

CREATE TABLE tasks (
  id bigserial primary key,
  title varchar(255) not null,
  detail text,
  status task_status not null,
  task_group_id bigint not null references task_groups(id),
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);

CREATE TABLE task_assignees (
  id bigserial primary key,
  task_id bigint not null references tasks(id),
  user_id bigint not null references users(id)
);
