startDB:
    docker run --name=postgres-v2 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:15.2-alpine
startCompile:
    CompileDaemon --command="./sql-api"
DockerRun:
    docker run -d -p 4300:4200 sql-api
DockerBuild:
    docker build -t sql-api .
DockerRemoveImage:
    docker rmi sql-api
DockerRemoveContainer:
    docker kill sql-api
    docker rm sql-api
DockerListPs:
    docker ps -a
DockerListImages:
    docker images   
.PHONY: pullPostgres startDB DockerRun DockerBuild DockerListPs DockerListImages DockerRemoveImage DockerRemoveContainer startCompile
Collapse















