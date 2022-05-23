import React, { useState } from 'react';
import { useQuery } from 'react-query';
import { useNavigate, useParams } from 'react-router-dom';
import styled from 'styled-components';
import { getInventory, getInventoryWarehouses } from '../api/InventoryAPI';
import InventoryForm from '../components/InventoryForm';
import ModalContainer from '../components/ModalContainer';
import WarehouseList from '../components/WarehouseList';
import Header from '../elements/Header';
import Main from '../elements/Main';
import useDeleteInventory from '../hooks/useDeleteInventory';

const InventoryPage = () => {
	const [updateModal, setUpdateModal] = useState(false);
	const [deleteModal, setDeleteModal] = useState(false);

	const { id } = useParams();
	const navigate = useNavigate();
	const { mutate: deleteInventory } = useDeleteInventory();
	const { data: inventory } = useQuery(
		['inventory', id],
		() => getInventory(id),
		{
			staleTime: 60 * 1000,
		}
	);
	const { data: warehouses } = useQuery(
		['inventory', 'warehouse', id],
		() => getInventoryWarehouses(id),
		{
			staleTime: 60 * 1000,
		}
	);

	const handleDelete = () => {
		deleteInventory(id);
		navigate('/', { replace: true });
	};

	const openUpdateModal = () => {
		setUpdateModal(true);
	};
	const openDeleteModal = () => {
		setDeleteModal(true);
	};
	const closeUpdateModal = () => {
		setUpdateModal(false);
	};
	const closeDeleteModal = () => {
		setDeleteModal(false);
	};

	return (
		<Main>
			<Header>Inventory</Header>
			<Info>
				<Name>{inventory && inventory.name}</Name>
				<Price>${inventory && inventory.price}</Price>
			</Info>
			<Tools>
				<Button onClick={openUpdateModal} background='rgb(59 130 246)'>
					Update
				</Button>
				<Button onClick={openDeleteModal}>Delete</Button>
			</Tools>
			<Header>Warehouses</Header>
			<WarehouseList items={warehouses} />
			<ModalContainer isOpen={updateModal} onClose={closeUpdateModal}>
				<InventoryForm value={inventory} onSubmit={closeUpdateModal} />
			</ModalContainer>
			<ModalContainer isOpen={deleteModal} onClose={closeDeleteModal}>
				<Container>
					<Header>Are you sure?</Header>
					<Button onClick={handleDelete}>Delete</Button>
				</Container>
			</ModalContainer>
		</Main>
	);
};

const Info = styled.div`
	display: flex;
	align-items: center;
	gap: 1rem;
	padding: 0.5rem 0;
`;
const Name = styled.h3`
	font-size: var(--text-xl);
	font-weight: var(--font-bold);
	width: min-content;
	line-height: 1.1;
	padding: 1rem;
	margin-bottom: 0.5rem;
	background-color: var(--color-gray-100);
	border-radius: var(--rounded-full);
`;
const Price = styled.div`
	font-family: var(--font-secondary);
	font-size: var(--text-lg);
	font-weight: var(--font-bold);
`;
const Tools = styled.div`
	display: flex;
	gap: 0.5rem;
`;
const Button = styled.button`
	font-family: var(--font-primary);
	font-weight: var(--font-bold);
	padding: 1rem;
	transition: var(--transition);
	color: var(--color-light);
	background-color: ${(props) =>
		props.background ? props.background : 'rgb(234 88 12)'};
	border-radius: var(--rounded-xl);

	cursor: pointer;

	:hover:enabled {
		box-shadow: var(--shadow-md);
	}

	:active:enabled {
		box-shadow: var(--shadow);
	}
`;

const Container = styled.div`
	display: flex;
	flex-direction: column;
	align-items: center;
	gap: 0.5rem;
`;

export default InventoryPage;
