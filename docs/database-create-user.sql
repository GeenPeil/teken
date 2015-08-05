
CREATE ROLE pechtold PASSWORD 'md55279b99979be20e5f546f4ba96239112' NOSUPERUSER NOCREATEDB NOCREATEROLE INHERIT LOGIN;

GRANT CONNECT ON DATABASE geenpeil TO pechtold;

GRANT SELECT, INSERT ON ALL TABLES IN SCHEMA public TO pechtold;

# edit pg_hba.conf to allow md5 auth for local unix socket connections
