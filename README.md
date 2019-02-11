# Install PostgreSQL

Install the latest Postgree official docker image
> docker pull postgres

Create the Postgres folder
> mkdir -p $HOME/docker/volumes/postgres

# Run the PostgreSQL

Following [this blog post.](https://hackernoon.com/dont-install-postgres-docker-pull-postgres-bee20e200198)  

> docker run --rm --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres

# Testing the Postgres server running
> psql -h localhost -U postgres -d postgres


# Useful links

* [the-myth-about-golang-frameworks-and-external-libraries](https://hackernoon.com/the-myth-about-golang-frameworks-and-external-libraries-93cb4b7da50f)
* [what-ive-learned-in-7-days-of-go](https://dev.to/rhymes/what-ive-learned-in-7-days-of-go-5e3)

# Idea

In this project I'm going to update the [Todo app](https://github.com/scotch-io/go-echo-vue-single-page-app) blogged in the [Scotch.io](https://scotch.io/tutorials/create-a-single-page-app-with-go-echo-and-vue) site. My target is to migrate the database to Postgress using [go-pg](https://github.com/go-pg/pg) from the SQLite. You may find the original app in the initial commit.
