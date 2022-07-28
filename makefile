BINARY_NAME=main.out

SCRIPTS=client/scripts
STYLES=client/styles
TEMPLATES=server/templates

build:
	tsc ${SCRIPTS}/*.ts
	sass ${STYLES}/:${STYLES}/
	sass ${TEMPLATES}/:${TEMPLATES}/
	go build -o ${BINARY_NAME} main.go

clean:
	go clean
	rm ${SCRIPTS}/*.js
	rm ${STYLES}/*.css
	rm ${STYLES}/*.css.map
	rm ${TEMPLATES}/**/*.css
	rm ${TEMPLATES}/**/*.css.map
	rm ${BINARY_NAME}

run:
	tsc ${SCRIPTS}/*.ts
	sass ${STYLES}/:${STYLES}/
	sass ${TEMPLATES}/:${TEMPLATES}/
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

test:
	go test -v ./resources/msword
	go test -v ./server/
	go test -v ./server/docx/
	go test -v ./server/html/
