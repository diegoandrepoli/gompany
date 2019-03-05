/**
 * Company database setup
 */

-- create schema
create schema collector;

-- create table
create table if not exists collector.companies(
	id serial,
    name text not null,
    zip char(10) not null,
    website text
);
