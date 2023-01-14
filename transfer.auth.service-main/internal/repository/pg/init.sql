CREATE TABLE users (
                       id UUID DEFAULT gen_random_uuid (),
                       first_name VARCHAR (255),
                       last_name VARCHAR (255),
                       email VARCHAR (255) NOT NULL UNIQUE ,
                       phone VARCHAR (255) NOT NULL UNIQUE,
                       сountry_code VARCHAR (255) NOT NULL,
                       birthdate VARCHAR (255) NOT NULL,
                       password VARCHAR (255) NOT NULL,
                       blocked boolean NOT NULL default false,
                       image VARCHAR (255),
                       created_at  timestamp without time zone default (now() at time zone 'utc'),
                       updated_at  timestamp without time zone default (now() at time zone 'utc')
);
select * from users