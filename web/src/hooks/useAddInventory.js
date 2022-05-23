import { useQueryClient, useMutation } from 'react-query';
import { postInventory } from '../api/InventoryAPI';

const useAddInventory = () => {
	const queryClient = useQueryClient();
	return useMutation((inventory) => postInventory(inventory), {
		onSuccess: (newInventory) => {
			const inventories = queryClient.getQueryData('inventories');
			queryClient.setQueryData('inventories', [
				...inventories,
				newInventory,
			]);
		},
	});
};

export default useAddInventory;
