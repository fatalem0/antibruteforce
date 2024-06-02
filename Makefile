.PHONY: unit
CU = "docker-compose.unit.yaml"

unit_test:
	go test -v -count=1 ./internal/usecase/ipusecase

up:
	docker-compose up  --build -d
psql:
	docker exec -it postgres psql -U postgres -d force
test:
	test_status=0;\
	docker-compose -f docker-compose.test.yaml up --build -d;\
	docker-compose -f docker-compose.test.yaml run integration_test go test -v -count=1 ./...  || test_status=$$?;\
	docker-compose -f docker-compose.test.yaml down; echo "status="$$test_status;exit $$test_status;

unitd:
	test_status=0;\
	docker-compose -f $(CU) up --build -d;\
	docker-compose -f $(CU) run unit go test -v -count=1 ./internal/usecase/ipusecase || test_status=$$?;\
	docker-compose -f $(CU) down; echo "status="$$test_status;exit $$test_status;


