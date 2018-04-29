# Poorbox
Archive and streaming platform

## Installing
Not a whole lot to see at the moment,
but

    go get github.com/yourfin/poorbox

is a good start.

### Database
Docker needs to be installed to do this the easy way, but then you can just run `pg-docker-dev.sh`. Otherwise install postgres manually somewhere and get it running

After that, create a database named poorbox on the running postgres instance. This can be done with the command `CREATE DATABASE POORBOX` on the postgres command line.

Following this, run `poorbox db init` to create the requisite tables.
