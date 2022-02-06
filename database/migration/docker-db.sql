create table resumes (
    id serial primary key,
    total_income real,
    total_outcome real,
    balance real
);

create table outcomes_cat (
	id integer primary key,
	category varchar NOT NULL UNIQUE
);

create table incomes (
    id serial primary key,
    describe varchar(255) NOT NULL check (describe != ''),
    value real NOT NULL check (value > 0),
    date varchar(11) NOT NULL check (date != ''),
    day integer NOT NULL,
    month integer NOT NULL,
    year integer NOT NULL,
  	constraint duplicate_income unique (describe, month)
);

create table outcomes (
    id serial primary key,
    describe varchar(255) NOT NULL check (describe != ''),
    value real NOT NULL check (value > 0),
    date varchar(11) NOT NULL check (date != ''),
	category varchar(25) NOT NULL check (category != '') references outcomes_cat(category), 
    day integer NOT NULL,
    month integer NOT NULL,
    year integer NOT NULL,
    constraint duplicate_outcome unique (describe, month)
);

create table users (
    id serial primary key,
    username varchar(255) UNIQUE NOT NULL check (username != ''),
    passwordHash varchar(255) NOT NULL check (passwordHash != ''),
    isDisable boolean
);

insert into outcomes_cat (id, category) values 
	(1, 'Alimentação'),
	(2, 'Saúde'),
	(3, 'Moradia'),
	(4, 'Transporte'),
	(5, 'Educação'),
	(6, 'Lazer'),
	(7, 'Imprevistos'),
	(8, 'Outras');