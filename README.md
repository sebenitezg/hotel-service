# Microservice - Hotel Service

# Requirements
    - Golang
    - Docker
    - Just
    - DBMate

# Setup
1. Setup Postgres DB:
```
just db-setup
```

2. Execute migrations:
```
just db up
```

3. Initialize configurations:
```
just init
```

4. Execute the application:
```
just run
```