.PHONY: up
up:
	@docker-compose up -d
	@echo "UP"

.PHONY: stop
stop:
	@docker-compose stop
	@echo "STOP"

.PHONY: down
down:
	@docker-compose down --volumes --remove-orphans
	@echo "DOWN"

.PHONY: reset
reset: down up
