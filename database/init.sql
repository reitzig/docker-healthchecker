CREATE TABLE data(
   id serial PRIMARY KEY,
   content VARCHAR (50) NOT NULL
);

INSERT INTO data (content)
VALUES
   ('hello'), 
   ('pretty'),
   ('world');
