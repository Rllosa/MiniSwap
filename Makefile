# localchain Makefile

# Define ANSI escape codes for colors
GREEN=\033[0;32m
YELLOW=\033[0;33m
NC=\033[0m # No Color


ENV ?= DEV

# Define a target that combines the above targets
all: go

deployment:
	@osascript deploy_contract.scpt

run: 
	@go run -ldflags "-X main.ENV=$(ENV)" main.go  deploy.go utils.go setup.go setEnvFiles.go interact.go

go:
	@echo ""
	@echo "$(YELLOW)Starting LocalChain...$(NC)"
	@make deployment
	@echo "$(GREEN)LocalChain started successfully!$(NC)"
	@sleep 2
	@echo "$(YELLOW)Deploying smart contracts on the localChain...$(NC)"
	@make run
	@echo "$(GREEN)Contracts deployed successfully!$(NC)"

.PHONY: start_ganache deployment go