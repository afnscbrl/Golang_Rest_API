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
	category varchar NOT NULL check (category != '') references outcomes_cat(category), 
    month integer NOT NULL default (extract(MONTH FROM CURRENT_DATE)),
    constraint duplicate_outcome unique (describe, month)
);

create table outcomes_cat (
	id integer primary key,
	category varchar NOT NULL UNIQUE
)

insert into outcomes_cat (id, category) values 
	(1, 'Alimentação'),
	(2, 'Saúde'),
	(3, 'Moradia'),
	(4, 'Transporte'),
	(5, 'Educação'),
	(6, 'Lazer'),
	(7, 'Imprevistos'),
	(8, 'Outras');