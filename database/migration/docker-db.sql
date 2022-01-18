create table resumes (
    id serial primary key,
    total_income real,
    total_outcome real,
    balance real
);

create table incomes (
    id serial primary key,
    describe varchar NOT NULL check (describe != ''),
    value real NOT NULL check (value > 0),
    date varchar NOT NULL check (date != ''),
  	month integer not null default (extract(MONTH FROM CURRENT_DATE)),
  	constraint duplicate_income unique (describe, month)
);

create table outcomes (
    id serial primary key,
    describe varchar NOT NULL check (describe != ''),
    value real NOT NULL check (value > 0),
    date varchar NOT NULL check (date != ''),
    month integer NOT NULL default (extract(MONTH FROM CURRENT_DATE)),
    constraint duplicate_outcome unique (describe, month)
);
