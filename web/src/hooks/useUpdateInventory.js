import { useMutation, useQueryClient } from 'react-query';
import { updateInventory } from '../api/InventoryAPI';

const useUpdateInventory = () => {
	const queryClient = useQueryClient();
	return useMutation(({ id, inventory }) => updateInventory(id, inventory), {
		onSuccess: (newInventory) => {
			const inventories = queryClient.getQueryData('inventories');
			queryClient.setQueryData(
				'inventories',
				updateInventories(inventories, newInventory)
			);
			queryClient.setQueryData(
				['inventory', newInventory._id],
				newInventory
			);
		},
	});
};

const updateInventories = (inventories, inventory) => {
	if (!inventory) return inventories;
	if (!inventories || inventories.length === 0) return [inventory];

	const index = inventories.findIndex((item) => item._id === inventory._id);
	if (index === -1) {
		inventories.unshift(inventory);
	} else {
		inventories[index] = inventory;
	}

	return inventories;
};

export default useUpdateInventory;
