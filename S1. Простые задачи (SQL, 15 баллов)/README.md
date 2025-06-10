# Простые задачи

Напишите запрос к базе данных, который возвращает все задачи, которые были решены не менее чем двумя пользователями. Найденные задачи следует отсортировать по id.

Внимательно ознакомьтесь с примерами вывода. Ваш запрос должен иметь в точности такой же вывод на примерах.

Схема базы данных содержит четыре таблицы:

* users — пользователи системы (описываются двумя полями: id и name),
* contests — контесты в системе (описываются двумя полями: id и name),
* problems — задачи в системе, каждая задача принадлежит одному контесту (описываются тремя полями: id, contest_id и code, где code — это кодовое короткое название задачи),
* submissions — отосланные попытки решения задач, каждая попытка принадлежит одной задаче и одному пользователю (описываются 5 полями: id, user_id, problem_id, success и submitted_at, где success — это булевское значение была ли попытка успешной и submitted_at — дата-время, когда попытка была совершена).
Таким образом, contests и problems находятся в отношении «один ко многим», submissions и users находятся в отношении «многие к одному», submissions и problems находятся в отношении «многие к одному».

Изучите входные данные примера, чтобы подробно ознакомиться со схемой базы данных. Диаграмма ниже иллюстрирует схему базы данных.

![img.png](img.png)

#### Входные данные
Входными данными в этой задаче является дамп базы данных. Вам он может быть полезен для ознакомления с состоянием базы данных для конкретного теста. В качестве решения вы должны отправить один SQL-запрос.

#### Выходные данные
Ваш SQL-запрос должен вывести все подходящие задачи в порядке возрастания их id.

Внимательно ознакомьтесь с примерами вывода. Ваш запрос должен иметь в точности такой же вывод на примерах.

#### Примеры

#### Входные данные
```sql
create table users (
  id bigint primary key,
  name varchar not null
);

create table contests (
  id bigint primary key,
  name varchar not null
);

create table problems (
  id bigint primary key,
  contest_id bigint,
  code varchar not null,
  constraint fk_problems_contest_id foreign key (contest_id) references contests (id)
);

create unique index on problems (contest_id, code);

create table submissions (
  id bigint primary key,
  user_id bigint,
  problem_id bigint,
  success boolean not null,
  submitted_at timestamp not null,
  constraint fk_submissions_user_id foreign key (user_id) references users (id),
  constraint fk_submissions_problem_id foreign key (problem_id) references problems (id)
);

insert into users
values (1, 'Marie Curie'),
       (2, 'Stephen Hawking'),
       (3, 'Ada Lovelace'),
       (4, 'Albert Einstein'),
       (5, 'Archimedes');

insert into contests
values (1, 'Sandbox-Juniors'),
       (2, 'Sandbox-Seniors'),
       (3, 'Contest-Juniors'),
       (4, 'Contest-Seniors');

insert into problems
values (1, 1, 'A'),
       (2, 2, 'A'),
       (3, 3, 'A'),
       (4, 3, 'B'),
       (5, 4, 'A'),
       (6, 4, 'B');

insert into submissions
values (1, 2, 2, false, '2023-02-05 11:01:00'),
       (2, 2, 2, true, '2023-02-05 11:02:00'),
       (3, 2, 6, true, '2023-02-05 11:03:01'),
       (4, 2, 1, true, '2023-02-05 11:04:00'),
       (5, 2, 1, true, '2023-02-05 11:05:00'),
       (6, 3, 6, true, '2023-02-05 11:06:00'),
       (17, 1, 6, true, '2023-02-05 11:03:00'),
       (8, 1, 2, true, '2023-02-05 11:08:00'),
       (9, 1, 1, false, '2023-02-05 11:09:00'),
       (10, 3, 1, false, '2023-02-05 11:10:00'),
       (11, 5, 5, false, '2023-02-05 11:11:00'),
       (13, 2, 6, true, '2023-02-05 11:03:00'),
       (14, 3, 6, false, '2023-02-05 11:05:59'),
       (15, 1, 6, true, '2023-02-05 11:04:00');
```

#### Выходные данные
```sql
 id | contest_id | code
----+------------+------
  2 |          2 | A
  6 |          4 | B
(2 rows)
```