import React, { useCallback, useEffect, useState } from 'react';
import { useQuery } from 'react-query';
import { useParams } from 'react-router-dom';
import styled from 'styled-components';
import { getInventories } from '../api/InventoryAPI';
import { getWarehouse, getWarehouseInventories } from '../api/WarehouseAPI';
import Location from '../components/Location';
import WarehouseInventoryList from '../components/WarehouseInventoryList';
import Header from '../elements/Header';
import Main from '../elements/Main';
import useAddWarehouseInventory from '../hooks/useAddWarehouseInventory';
import useDeleteWarehouseInventory from '../hooks/useDeleteWarehouseInventory';
import AddIcon from '../icons/AddIcon';
import CloseIcon from '../icons/CloseIcon';

const WarehousePage = () => {
	const { id } = useParams();
	const { data: warehouse = {} } = useQuery(
		['warehouse', id],
		() => getWarehouse(id),
		{
			staleTime: 60 * 1000,
		}
	);
	const { data: warehouseInventories } = useQuery(
		['warehouse-inventory', id],
		() => getWarehouseInventories(id),
		{
			staleTime: 60 * 1000,
		}
	);
	const { data: inventories } = useQuery(
		'inventories',
		() => getInventories(),
		{
			staleTime: 60 * 1000,
		}
	);

	const [availableInventories, setAvailableInventories] = useState([]);
	const updateAvailableInventories = useCallback(
		(warehouseInventories) => {
			setAvailableInventories(
				filterInventories(warehouseInventories, inventories)
			);
		},
		[inventories]
	);

	const { mutate: removeInventory } = useDeleteWarehouseInventory(
		id,
		updateAvailableInventories
	);
	const { mutate: addInventory } = useAddWarehouseInventory(
		warehouse,
		updateAvailableInventories
	);

	useEffect(() => {
		updateAvailableInventories(warehouseInventories);
	}, [warehouseInventories, inventories, updateAvailableInventories]);

	const handleRemoveInventory = (inventory) => {
		removeInventory(inventory._id);
	};

	const handleAddInventory = (inventory) => {
		addInventory(inventory);
	};

	return (
		<Main>
			<Header>Warehouse</Header>
			<Container>
				<Name>{warehouse && warehouse.name}</Name>
				<Location value={warehouse.location} />
			</Container>
			<ListContainer>
				<List>
					<Header>Warehouse Inventories</Header>
					<WarehouseInventoryList
						items={warehouseInventories}
						onClick={handleRemoveInventory}
						icon={
							<IconContainer color='var(--color-error)'>
								<CloseIcon />
							</IconContainer>
						}
					/>
				</List>
				<List>
					<Header>Available Inventories</Header>
					<WarehouseInventoryList
						items={availableInventories}
						onClick={handleAddInventory}
						icon={
							<IconContainer>
								<AddIcon />
							</IconContainer>
						}
					/>
				</List>
			</ListContainer>
		</Main>
	);
};

const filterInventories = (warehouseInventories, inventories) => {
	if (!warehouseInventories) return inventories || [];
	return inventories.filter(
		(a) => !warehouseInventories.some((b) => a._id === b._id)
	);
};

const Container = styled.div`
	padding: 2rem;
	border-radius: var(--rounded-3xl);
	box-shadow: var(--shadow);
`;
const Name = styled.h3`
	font-size: var(--text-4xl);
	font-weight: var(--font-bold);
	margin-bottom: 1rem;
`;
const ListContainer = styled.div`
	display: flex;
	gap: 2rem;
	margin-top: 1rem;
`;
const List = styled.div`
	flex-basis: 50%;
`;
const IconContainer = styled.div`
	position: absolute;
	z-index: 50;
	top: -0.5rem;
	right: -0.5rem;
	padding: 0.5rem;

	cursor: pointer;

	color: var(--color-light);
	background-color: ${(props) =>
		props.color ? props.color : 'var(--color-success)'};
	border-radius: var(--rounded-full);
	transition: var(--transition);

	:hover {
		filter: brightness(1.1);
	}
`;

export default WarehousePage;
