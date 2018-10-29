use prac;
alter table item add foreign key (user_id) references user (id);