# Set Up for Development

## Config

todo: update config explanation

## Spinning up locally

```
BUILD_ENV=development docker-compose up --build
```

If using the gogetter-cli:

```
gogetter run dev
```

## Usage

After the containers have been set up locally, the client can be accessed at [http://localhost:3000]. The server can be accessed at [http://localhost:3001].

## Shutting down

```
docker-compose down --remove-orphans
```

If using the gogetter-cli:

```
gogetter stop
```

OR

```
gogetter down
```

## Troubleshooting

- For node_module errors:
  1. delete the node modules folder
  2. cd into client
  3. run `npm install`
  4. spin up containers

## For deployment

todo update section

```
BUILD_ENV=production docker-compose up --build
```

---

# Architectural Information

## Client

- language: javascript
- framework: Vue3 with Vite
- architecture follows framework conventions
- store:
  - library: pinia

## Server

- language: Go
- server framework: Echo
- database: PostgreSQL
- ORM: [GORM] (https://gorm.io/)
- DI: [Fx]

### Style

- follow [Uber-Go Style Guide](https://github.com/uber-go/guide) wherever possible
- architecture is domain driven
- follow CLEAN & SOLID principles wherever feasible

### Architecture

- domain-driven
- dependency injection using Fx framework

### Folder Structure

To Do: explain directories

Common domain:

- Contains Models, Errors, etc. common between various domains

Schema domain:

- Contains database schema structs

## Database

- A local postgreSQL database included in docker-compose setup
- if using pgAdmin4 or datagrip to connect to the local postgres container, use `host:localhost`

---

# API Documentation

Todo: document API routes

---

### Git Conventions

Use [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/#summary)
