import { useMutation, useQueryClient } from 'react-query';
import { deleteInventory } from '../api/InventoryAPI';

const useDeleteInventory = () => {
	const queryClient = useQueryClient();
	return useMutation((id) => deleteInventory(id), {
		onSuccess: (id) => {
			const inventories = queryClient.getQueryData('inventories');
			queryClient.setQueryData(
				'inventories',
				removeInventory(inventories, id)
			);
		},
	});
};

const removeInventory = (inventories, id) => {
	const index = inventories.findIndex((item) => item._id === id);
	if (index > -1) {
		inventories.splice(index, 1);
	}
	return inventories;
};

export default useDeleteInventory;
