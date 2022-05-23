import React, { useState } from 'react';
import { useQuery } from 'react-query';
import { getInventories } from '../api/InventoryAPI';
import { getWarehouses } from '../api/WarehouseAPI';
import AddButton from '../components/AddButton';
import InventoryForm from '../components/InventoryForm';
import InventoryList from '../components/InventoryList';
import ModalContainer from '../components/ModalContainer';
import WarehouseForm from '../components/WarehouseForm';
import WarehouseList from '../components/WarehouseList';
import Header from '../elements/Header';
import Main from '../elements/Main';

const HomePage = () => {
	const [inventoryModalIsOpen, setInventoryModalIsOpen] = useState(false);
	const [warehouseModalIsOpen, setWarehouseModalIsOpen] = useState(false);

	const { data: inventories } = useQuery(
		'inventories',
		() => getInventories(),
		{
			staleTime: 60 * 1000,
		}
	);
	const { data: warehouses } = useQuery('warehouses', () => getWarehouses(), {
		staleTime: 60 * 1000,
	});

	const handleAddInventory = () => {
		setInventoryModalIsOpen(true);
	};

	const handleAddWarehouse = () => {
		setWarehouseModalIsOpen(true);
	};

	const closeInventoryModal = () => {
		setInventoryModalIsOpen(false);
	};

	const closeWarehouseModal = () => {
		setWarehouseModalIsOpen(false);
	};

	return (
		<Main>
			<Header>Inventories</Header>
			<InventoryList items={inventories}>
				<AddButton onClick={handleAddInventory} />
			</InventoryList>
			<Header>Warehouses</Header>
			<WarehouseList items={warehouses}>
				<AddButton onClick={handleAddWarehouse} />
			</WarehouseList>
			<ModalContainer
				isOpen={inventoryModalIsOpen}
				onClose={closeInventoryModal}
			>
				<InventoryForm onSubmit={closeInventoryModal} />
			</ModalContainer>
			<ModalContainer
				isOpen={warehouseModalIsOpen}
				onClose={closeWarehouseModal}
			>
				<WarehouseForm onSubmit={closeWarehouseModal} />
			</ModalContainer>
		</Main>
	);
};

export default HomePage;
