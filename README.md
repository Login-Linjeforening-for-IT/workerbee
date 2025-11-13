# Workerbee API
This is the main API for Beehive and Queenbee applications. It is built using Go and Gin framework, and uses PostgreSQL as the database. It has both protected and public endpoints along with comprehensive documentation. Its endpoints are documented using swagger and can be accessed a `/api/v2/docs` once the application is running. The API is capable of both talking to a PostgreSQL database and a DigitalOcean Spaces instance for file storage, this is used to store images.

# Beehive Database
This is the database currently used by Login for our beehive production environment. It contains a `db/init.sql` file with the structure of the database, and a `db/dummydata.sql` file which contains the same structure but has also been populated with testdata.

## Running the application
### Production:
The production database is set up using the `init.sql` file.  
Start it with `docker compose up --build` \
This will start both the database and the API server.

### Development:
The development database is set up using the `dummydata.sql` which means its automatically populated with test data.  
Remember to set the `INIT_SQL_FILE` variabel to `./db/dummydata.sql`  
Start it with `docker compose up --build`

### Environment Variables
The application uses the following environment variables. These can be set in a `.env` file in the root directory.

| Variable                 | Default Value   | Description                                      |
|--------------------------|-----------------|--------------------------------------------------|
| POSTGRES_PASSWORD        | (required)      | Password for the PostgreSQL database.            |
| INIT_SQL_FILE            | `./db/init.sql` | Path to the SQL file to initialize the database. |
| POSTGRES_USER            | (required)      | Username for the PostgreSQL database.            |
| POSTGRES_DB              | (required)      | Name of the PostgreSQL database.                 |
| POSTGRES_HOST            | (required)      | Hostname for the PostgreSQL database.            |
| POSTGRES_PORT            | (optional)      | Port for the PostgreSQL database.                |
| PORT                     | (optional)      | Port for the API server to listen on.            |
| DO_URL                   | (required)      | DigitalOcean Spaces endpoint URL.                |
| DO_SECRET_ACCESS_KEY     | (required)      | DigitalOcean Spaces secret access key.           |
| DO_ACCESS_KEY_ID         | (required)      | DigitalOcean Spaces access key ID.               |