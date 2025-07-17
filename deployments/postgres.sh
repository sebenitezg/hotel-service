#!/bin/bash
# This script deploys a PostgreSQL database for the hotel service.

docker run --name hotel-service-db \
    -e POSTGRES_PASSWORD=secretpassword \
    -e POSTGRES_USER=developer \
    -e POSTGRES_DB=hotel_service \
    -p 5431:5432 \
    -d postgres
