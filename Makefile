-include .env

# Força o shell usado pelo make a ser bash.
# Por padrão, ele usa o sh, mas como sh está depreciado,
# diferentes distribuições linux usa, sh como um alias para
# algum outro shell, por exemplo, ubuntu usa dash e arch usa bash.
# Forçar o shell para bash permite que o make seja mais consistente
# entre distribuições linux diferentes
SHELL := /bin/bash

run-acceptance-tests:
	docker compose --project-directory acceptance-tests up \
		--build --quiet-pull \
		--no-log-prefix --exit-code-from acceptance-tests \
		acceptance-tests
	docker compose --project-directory acceptance-tests down