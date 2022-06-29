FRONTEND := ./frontend
BACKEND := ./backend

define execute_go_api
	cd backend && go run main.go &
endef

define execute_frontend
	cd frontend && yarn serve &
endef

define return_to_base_dir
	cd ..
endef

define run_docker_environment
	docker compose up --build
endef

dev_local: dev_frontend return_to_base
	$(call execute_go_api)

dev_frontend:
	$(call execute_frontend)

return_to_base:
	$(call return_to_base_dir)

dev_docker:
	$(call run_docker_environment)