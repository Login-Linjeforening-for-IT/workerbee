#!/bin/bash

if [ "$LOAD_DUMMY_DATA" = "true" ]; then
    echo "Loading dummy data..."
    psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f /dummydata.sql
    echo "Dummy data loaded successfully."
else
    echo "LOAD_DUMMY_DATA is not set to 'true'. Skipping dummy data."
fi