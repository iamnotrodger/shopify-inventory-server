import { buildQueryString } from './query';
import RequestError from './RequestError';

export const getWarehouses = async (query) => {
	const queryString = buildQueryString(query);
	const response = await fetch(`/api/warehouse${queryString}`);

	if (!response.ok) throw await RequestError.parseResponse(response);

	const inventories = await response.json();
	return inventories;
};

export const getWarehouse = async (id) => {
	const response = await fetch(`/api/warehouse/${id}`);

	if (!response.ok) throw await RequestError.parseResponse(response);

	const warehouse = await response.json();
	return warehouse;
};

export const getWarehouseInventories = async (id, query) => {
	const queryString = buildQueryString(query);
	const response = await fetch(
		`/api/warehouse/${id}/inventory${queryString}`
	);

	if (!response.ok) throw await RequestError.parseResponse(response);

	const warehouses = await response.json();
	return warehouses;
};

export const postWarehouse = async (warehouse) => {
	const response = await fetch('/api/warehouse', {
		method: 'POST',
		mode: 'cors',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(warehouse),
	});

	if (!response.ok) throw await RequestError.parseResponse(response);

	const warehouseRes = await response.json();
	return warehouseRes;
};

export const deleteWarehouse = async (id) => {
	const response = await fetch(`/api/warehouse/${id}`, {
		method: 'DELETE',
		mode: 'cors',
	});

	if (!response.ok) throw await RequestError.parseResponse(response);
};

export const postInventoryToWarehouse = async (warehouseID, inventoryID) => {
	const response = await fetch(
		`/api/warehouse/${warehouseID}/inventory/${inventoryID}`,
		{
			method: 'POST',
			mode: 'cors',
		}
	);

	if (!response.ok) throw await RequestError.parseResponse(response);
};

export const deleteInventoryFromWarehouse = async (
	warehouseID,
	inventoryID
) => {
	const response = await fetch(
		`/api/warehouse/${warehouseID}/inventory/${inventoryID}`,
		{
			method: 'DELETE',
			mode: 'cors',
		}
	);

	if (!response.ok) throw await RequestError.parseResponse(response);
};
