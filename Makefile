

swagger-generate: ## generate swagger
	swagger generate server -A TodoList -f swagger.yml

# Usage: make migration table=your_table_name
migration: ## create migration
	if [ -z $${table} ]; then echo "table is required"; exit 1; fi
	echo "CREATE TABLE IF NOT EXISTS ${table}" > ./migrations/$$(date +%Y%m%d%H%M%S)_$${table}.sql
