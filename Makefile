GCC=go
GCMD=run
GPATH=main.go
# GO_VERSION=1.14.1

run:
	# make build
	$(GCC) $(GCMD) $(GPATH)
build:
	make build_db

build_db:
	rm pkg/db/db_struct.go
	dgw postgres://postgres:password@localhost/testdb?sslmode=disable --schema=public --package=db --output=db_struct.go
	mv db_struct.go pkg/db/