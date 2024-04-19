curl-root:
	@ curl -sS -r GET http://localhost:8080

curl-param:
	@ curl -sS -r GET http://localhost:8080/v1/param/200

goose-up:
	@ ./sql/scripts/goose-up.sh

goose-down:
	@ ./sql/scripts/goose-down.sh

goose-reset:
	@ ./sql/scripts/goose-reset.sh

sqlc:
	@ sqlc generate
