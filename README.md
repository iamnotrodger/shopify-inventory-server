# Shopfiy-Inventory-Server

## Chosen Feature

Ability to create warehouses/locations and assign inventory to specific locations

## CRUD Operations

| METHOD | URI | Description |
| ------- | -------  | ------- |
| GET | / | React web application |
| GET | /api/inventory | Get multiple inventories |
| POST | /api/inventory | Create inventory item |
| GET | /api/inventory/:inventoryID | Get inventory with `inventoryID` |
| PUT | /api/inventory/:inventoryID | Edit entire inventory |
| PATCH | /api/inventory/:inventoryID | Partially edit inventory |
| DELETE | /api/inventory/:inventoryID | Delete inventory with `inventoryID` |
| GET | /api/inventory/:inventoryID/warehouse | Get inventory assigned warehouses |
| GET | /api/warehouse | Get multiple warehouses |
| POST | /api/warehouse | Create warehouse |
| GET | /api/warehouse/:warehouseID | Get warehouse with `warehouseID` |
| GET | /api/warehouse/:warehouseID/inventory | Get warehouse inventories |
| POST | /api/warehouse/:warehouseID/inventory/:inventoryID | Assign inventory to warehouse |
| DELETE | /api/warehouse/:warehouseID/inventory/:inventoryID | Remove inventory to warehouse |

For operations that gets multiple items, all URI endpoints take in query parameters `limit` and `skip`.

Examples:

```bash
curl --location --request GET 'localhost:8080/api/inventory?limit=3&skip=1'
curl --location --request GET 'localhost:8080/api/inventory/60e0c4aeffdd3e5211a78a32/warehouse?limit=3&skip=1'
curl --location --request GET 'localhost:8080/api/warehouse?limit=3&skip=1'
curl --location --request POST 'localhost:8080/api/inventory' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Inventory-1",
    "price": 10.10
}'
```

## Models

### Inventory

```json
{
	"_id": <string>,
	"name": <string>,
	"price": <float>,
}
```

### Warehouse

```json
{
	"_id": <string>,
	"name": <string>,
	"location": {
		"street": <string>,
		"city": <string>,
		"province": <string>,
		"country": <string>
	}
}
```