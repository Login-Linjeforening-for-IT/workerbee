# Beehive Database
This is the database currently used by Login for our beehive production environment. It contains a `db/init.sql` file with the structure of the database, and a `db/dummydata.sql` file which contains the same structure but has also been populated with testdata.

## Running the application
### Production:
The production database is set up using the `init.sql` file.  
Start it with `docker compose up --build`

### Development:
The development database is set up using the `dummydata.sql` which means its automatically populated with test data.  
Remember to set the `INIT_SQL_FILE` variabel to `./db/dummydata.sql`  
Start it with `docker compose up --build`

### Environment Variables
The application uses the following environment variables. These can be set in a `.env` file in the root directory.

| Variable          | Default Value   | Description                                      |
|-------------------|-----------------|--------------------------------------------------|
| POSTGRES_PASSWORD | (required)      | Password for the PostgreSQL database.            |
| INIT_SQL_FILE     | `./db/init.sql` | Path to the SQL file to initialize the database. |