# The Honest Truth

## Purpose

The purpose of this application is to provide users with the ability to annotate and reference portions of youtube videos interactively.

## Build Instructions

### UI API (backend)

```bash
$ cd ui-api
$ ./mvnw clean install springboot:run
```

### UI (Frontend)

```bash
$ cd honest-truth-ui
$ npm start
```

### Database

```bash
$ docker-compose up
```