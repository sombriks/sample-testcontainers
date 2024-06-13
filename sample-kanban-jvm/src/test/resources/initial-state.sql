-- board schema
create schema kanban;

-- meaningful statuses
create table kanban.status
(
    id             serial primary key,
    description    text unique not null,
    means_complete boolean              default false,
    created        timestamp   not null default current_timestamp
);

-- people that does things
create table kanban.person
(
    id      serial primary key,
    name    text      not null,
    created timestamp not null default current_timestamp
);

-- things being done
create table kanban.task
(
    id          serial primary key,
    status_id   integer   not null,
    description text      not null,
    created     timestamp not null default current_timestamp,
    foreign key (status_id) references kanban.status (id)
);

-- meaningful messages about tasks, from the people
create table kanban.message
(
    id        serial primary key,
    person_id integer   not null,
    task_id   integer   not null,
    content   text      not null,
    created   timestamp not null default current_timestamp,
    foreign key (person_id) references kanban.person (id),
    foreign key (task_id) references kanban.task (id)
);

-- who is working on what
create table kanban.task_person
(
    person_id integer   not null,
    task_id   integer   not null,
    created   timestamp not null default current_timestamp,
    foreign key (person_id) references kanban.person (id),
    foreign key (task_id) references kanban.task (id),
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

-- some tasks
insert into kanban.task(description, status_id)
values ('design',
        (select id
         from kanban.status
         where description = 'DOING')),
       ('data provision',
        (select id
         from kanban.status
         where description = 'TODO')),
       ('project initial setup',
        (select id
         from kanban.status
         where description = 'DONE')),
       ('feature listing',
        (select id
         from kanban.status
         where description = 'DONE')),
       ('server configuration',
        (select id
         from kanban.status
         where description = 'TODO'));

-- people working on tasks
insert into kanban.task_person(person_id, task_id)
values ((select id
         from kanban.person
         where name = 'Alice'),
        (select id
         from kanban.task
         where description = 'server configuration')),
       ((select id
         from kanban.person
         where name = 'Bob'),
        (select id
         from kanban.task
         where description = 'server configuration')),
       ((select id
         from kanban.person
         where name = 'Caesar'),
        (select id
         from kanban.task
         where description = 'feature listing'));

-- and a comment
insert into kanban.message(person_id, task_id, content)
values ((select id
         from kanban.person
         where name = 'Caesar'),
        (select id
         from kanban.task
         where description = 'feature listing'),
        'Need this ASAP');
