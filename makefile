BINARY_NAME=main.out

SCRIPTS=client/scripts
STYLES=client/styles
TEMPLATES=server/templates

build:
	tsc ${SCRIPTS}/*.ts
	sass ${STYLES}/:${STYLES}/
	sass ${TEMPLATES}/:${TEMPLATES}/
	go build -o ${BINARY_NAME} main.go

run:
	tsc ${SCRIPTS}/*.ts
	sass ${STYLES}/:${STYLES}/
	sass ${TEMPLATES}/:${TEMPLATES}/
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${SCRIPTS}/*.js
	rm ${STYLES}/*.css
	rm ${STYLES}/*.css.map
	rm ${TEMPLATES}/*.css
	rm ${TEMPLATES}/*.css.map
	rm ${BINARY_NAME}