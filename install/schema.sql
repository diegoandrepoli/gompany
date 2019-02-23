/**
 * Company database setup
 */

-- create database gompany
create database gompany;

-- create database companies in gompany database
create table if not exists gompany.companies(
	  id serial,
    name text not null,
    zip char(5) not null,
    website text
);
