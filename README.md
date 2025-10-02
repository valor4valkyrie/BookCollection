# Book Collection

## Description
This is a simple book collection app written in Go that tracks books that I have bought, sell or have borrowed to others.
Only a backend service is provided. Making this essentially a Restful CRUD API.

## Application Settings
Application settings are stored in the `properties.yaml` file. Located in the `properties` folder. The current settings are:
- db:
    - host:
    - port:
    - username:
    - password:

## Dockerfile
A Dockerfile is provided to build the application and have it run successfully. However, the application alone will not 
run completely as there is no database connection.