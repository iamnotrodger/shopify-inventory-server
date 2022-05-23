import { buildQueryString } from './query';
import RequestError from './RequestError';

export const getInventories = async (query) => {
	const queryString = buildQueryString(query);
	const response = await fetch(`/api/inventory${queryString}`);

	if (!response.ok) throw await RequestError.parseResponse(response);

	const inventories = await response.json();
	return inventories;
};

export const getInventory = async (id) => {
	const response = await fetch(`/api/inventory/${id}`);

	if (!response.ok) throw await RequestError.parseResponse(response);

	const inventory = await response.json();
	return inventory;
};

export const getInventoryWarehouses = async (id, query) => {
	const queryString = buildQueryString(query);
	const response = await fetch(
		`/api/inventory/${id}/warehouse${queryString}`
	);

	if (!response.ok) throw await RequestError.parseResponse(response);

	const warehouses = await response.json();
	return warehouses;
};

export const postInventory = async (inventory) => {
	const response = await fetch('/api/inventory', {
		method: 'POST',
		mode: 'cors',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(inventory),
	});

	if (!response.ok) throw await RequestError.parseResponse(response);

	const inventoryRes = await response.json();
	return inventoryRes;
};

export const deleteInventory = async (id) => {
	const response = await fetch(`/api/inventory/${id}`, {
		method: 'DELETE',
		mode: 'cors',
	});

	if (!response.ok) throw await RequestError.parseResponse(response);
	return id;
};

export const updateInventory = async (id, inventory) => {
	const response = await fetch(`/api/inventory/${id}`, {
		method: 'PUT',
		mode: 'cors',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(inventory),
	});

	if (!response.ok) throw await RequestError.parseResponse(response);

	const updatedInventory = await response.json();
	return updatedInventory;
};

export const patchInventory = async (id, inventory) => {
	const response = await fetch(`/api/inventory/${id}`, {
		method: 'PATCH',
		mode: 'cors',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(inventory),
	});

	if (!response.ok) throw await RequestError.parseResponse(response);

	const updatedInventory = await response.json();
	return updatedInventory;
};
