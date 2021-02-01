migrateup:
	migrate -path config/db/migrations -database "postgresql://docker:docker@db:5432/gis" -verbose up


migratedown:
	migrate -path config/db/migrations -database "postgresql://docker:docker@db:5432/gis" -verbose down
