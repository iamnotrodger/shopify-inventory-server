import { useMutation, useQueryClient } from 'react-query';
import { postInventoryToWarehouse } from '../api/WarehouseAPI';

const useAddWarehouseInventory = (warehouse, callback) => {
	const queryClient = useQueryClient();

	const addInventory = async (inventory) => {
		await postInventoryToWarehouse(warehouse._id, inventory._id);
		return inventory;
	};

	return useMutation(addInventory, {
		onSuccess: (inventory) => {
			const inventories =
				queryClient.getQueryData([
					'warehouse-inventory',
					warehouse._id,
				]) || [];
			const newInventories = [...inventories, inventory];
			queryClient.setQueryData(
				['warehouse-inventory', warehouse._id],
				newInventories
			);

			const warehouses =
				queryClient.getQueryData([
					'inventory-warehouses',
					inventory._id,
				]) || [];
			queryClient.setQueryData('inventory-warehouses', [
				...warehouses,
				warehouse,
			]);

			if (callback) callback(newInventories);
		},
	});
};

export default useAddWarehouseInventory;
