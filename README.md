# Install PostgreSQL

Install the latest Postgree official docker image
> docker pull postgres

Create the Postgres folder
> mkdir -p $HOME/docker/volumes/postgres

# Run the PostgreSQL

Following [this blog post.](https://hackernoon.com/dont-install-postgres-docker-pull-postgres-bee20e200198)  

> docker run --name pg-docker -e POSTGRES_USER=logan -e POSTGRES_PASSWORD=xmen -e POSTGRES_DB=marvel -d -p 5432:5432 postgres
 
# Run the app
> go run todo.go

# Testing the Postgres server running
> psql -h localhost -U logan -d marvel


# Useful links

* [the-myth-about-golang-frameworks-and-external-libraries](https://hackernoon.com/the-myth-about-golang-frameworks-and-external-libraries-93cb4b7da50f)    "
* [what-ive-learned-in-7-days-of-go](https://dev.to/rhymes/what-ive-learned-in-7-days-of-go-5e3)

# Idea

In this project I've updated the [Todo app](https://github.com/scotch-io/go-echo-vue-single-page-app) blogged in the [Scotch.io](https://scotch.io/tutorials/create-a-single-page-app-with-go-echo-and-vue) site. Modified app uses Postgres through [go-pg](https://github.com/go-pg/pg) instead of SQLite. I have written a blog post for this project [here](https://medium.com/p/15b2df362ecf).