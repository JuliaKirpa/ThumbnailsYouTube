# Microservice for downloading YT TN
___
## Contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)
* [Example](#example)

## General info
gRPC service for downloading thumbnails(preview of YouTube video) to sqlite database.
CLI Client for service: [CLI-Client](https://github.com/JuliaKirpa/ThumbnailsYouTube-CLI-Client-)

Service download Image with given URL and saves it to the db in the BLOB format.
CI implemented with GitHub Actions and contains 3 jobs:
- Golangci-lint
- Test
- Build

Service returns status and ID of image rows in db. If Image was downloaded before service returns status "already downloaded".
If service gets new url returns status "downloaded".

## Technologies
Project is created with:
* GO
* SQLite
* Docker
* gRPC
* GitHub Actions
* golangci-lint

## Setup

Сheck that you have free ports:

* 50051

if not, replace with yours in


* Dockerfile

```
    EXPOSE 50051
```

##in progress


