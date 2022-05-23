import { useMutation, useQueryClient } from 'react-query';
import { deleteInventoryFromWarehouse } from '../api/WarehouseAPI';

const useDeleteWarehouseInventory = (warehouseID, callback) => {
	const queryClient = useQueryClient();
	return useMutation(
		(inventoryID) => deleteInventoryFromWarehouse(warehouseID, inventoryID),
		{
			onSuccess: (inventoryID) => {
				const inventories =
					queryClient.getQueryData([
						'warehouse-inventory',
						warehouseID,
					]) || [];
				const newInventories = removeInventory(
					inventories,
					inventoryID
				);
				queryClient.setQueryData('warehouse-inventory', newInventories);

				const warehouses =
					queryClient.getQueryData([
						'inventory-warehouses',
						inventoryID,
					]) || [];
				queryClient.setQueryData(
					'inventory-warehouses',
					removeInventory(warehouses, warehouseID)
				);
				if (callback) callback(newInventories);
			},
		}
	);
};

const removeInventory = (inventories, id) => {
	const index = inventories.findIndex((item) => item._id === id);
	if (index > -1) {
		inventories.splice(index, 1);
	}
	return inventories;
};

export default useDeleteWarehouseInventory;
