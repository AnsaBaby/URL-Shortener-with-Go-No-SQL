CREATE KEYSPACE IF NOT EXISTS urlshortener
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};
