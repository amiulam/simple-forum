ALTER TABLE users
ADD username varchar(100) NOT NULL;

ALTEr TABLE users
ADD CONSTRAINT UNIQUE unique_username (username);