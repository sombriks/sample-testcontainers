-- board schema
create schema kanban;
-- meaningful statuses
create table kanban.status(
  id serial primary key,
  description text unique not null,
  means_complete boolean default false,
  created timestamp not null default current_timestamp
);
-- people that does things
create table kanban.person(
  id serial primary key,
  name text not null,
  created timestamp not null default current_timestamp
);
-- things being done
create table kanban.task(
  id serial primary key,
  status_id integer not null,
  description text not null,
  created timestamp not null default current_timestamp,
  foreign key (status_id) references kanban.status(id)
);
-- meaningful messages about tasks, from the people
create table kanban.message(
  id serial primary key,
  person_id integer not null,
  task_id integer not null,
  content text not null,
  created timestamp not null default current_timestamp,
  foreign key (person_id) references kanban.person(id),
  foreign key (task_id) references kanban.task(id)
);
-- who is working on what
create table kanban.task_person(
  person_id integer not null,
  task_id integer not null,
  created timestamp not null default current_timestamp,
  foreign key (person_id) references kanban.person(id),
  foreign key (task_id) references kanban.task(id),
  primary key (person_id, task_id)
);
--
-- for the sake of commodity, let's feed the database with something
--

-- statues for the board
insert into kanban.status(description, means_complete)
values ('TODO', false),
  ('DOING', false),
  ('DONE', true);
-- people to work
insert into kanban.person(name)
values ('Alice'),
  ('Bob'),
  ('Caesar'),
  ('David'),
  ('Edward');