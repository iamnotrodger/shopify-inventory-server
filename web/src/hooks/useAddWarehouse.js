import { useMutation, useQueryClient } from 'react-query';
import { postWarehouse } from '../api/WarehouseAPI';

const useAddWarehouse = () => {
	const queryClient = useQueryClient();
	return useMutation((warehouse) => postWarehouse(warehouse), {
		onSuccess: (newWarehouse) => {
			const warehouses = queryClient.getQueryData('warehouses');
			queryClient.setQueryData('warehouses', [
				...warehouses,
				newWarehouse,
			]);
		},
	});
};

export default useAddWarehouse;
