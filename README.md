# Forma Shared

Overview
--------------------------------------------------------------------------------

This application helps the trainers during their training sessions.

It allows to share files locally for people in training.
There is also a Question and Answer system, for mutual help between students. This avoids repetitive questions to the trainer.

For the trainer, there is a checkpoint system in order to know the progress of each student on an particular exercise.


Status
--------------------------------------------------------------------------------

This application is functional.

Currently being optimized.

Also implementing the backoffice part, because currently, some things have to be changed directly in the database. (i.e. : if a user is admin or prof)


Client Setup
--------------------------------------------------------------------------------

### Build
* Go 1.5 (or higher) is required.
* This project builds and runs on recent versions of Windows, Linux, and Mac OS X.

> Download application
```
git clone https://github.com/kameleon83/forma_shared.git

```
> Enter into folder
```
cd forma_shared

```
> Get dependencies
```
cd forma_shared
go get

```
> Start application

```
go run main.go

```

#### Build Windows

```
go build -ldflags "-X main.buildstamp=$((get-date).tostring('d-MM-yyyy_H:mm:ss')) -X main.githash=$(git rev-parse HEAD) -X main.version=0.0.0 "

```

#### Build Linux

```
go build -ldflags "-X main.buildstamp=`date '+%d-%m-%Y_%H:%M:%S'` -X main.githash=`git rev-parse HEAD` -X main.version=0.0.0"

```


