deploy: 
	@npm install --prefix web/
	@npm run build --prefix web/
	@go mod download
	@go build -v -o . ./...
	@./shopify-inventory-server

run: ./cmd/shopify-inventory-server/main.go 
	@go run ./cmd/shopify-inventory-server/main.go --dev

run-web: ./web/package.json ./web/package-lock.json 
	@npm install --prefix web/
	@npm start --prefix web/

build-web: ./web/package.json ./web/package-lock.json
	@npm run build --prefix web/

seed: ./seed/inventories.json ./seed/warehouse.json
	@mongo shopify-inventory --eval "db.inventory.drop()"
	@mongo shopify-inventory --eval "db.createCollection('inventory')"
	@mongoimport --db shopify-inventory --collection inventory --file ./seed/inventories.json --jsonArray

	@mongo shopify-inventory --eval "db.warehouse.drop()"
	@mongo shopify-inventory --eval "db.createCollection('warehouse')"
	@mongoimport --db shopify-inventory --collection warehouse --file ./seed/warehouse.json --jsonArray