#!/bin/bash

docker run \
	--mount source=dbfiles,target=/var/lib/postgresql/data \
	--name pg-test \
  -e POSTGRES_PASSWORD=my_S3cur3-p4s5wrd \
	-p 5432:5432 \
	-d postgres

