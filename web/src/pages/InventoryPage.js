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
		['inventory-warehouses', id],
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
				<Tools>
					<Button
						onClick={openUpdateModal}
						background='rgb(59 130 246)'
					>
						Update
					</Button>
					<Button onClick={openDeleteModal}>Delete</Button>
				</Tools>
			</Info>
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
	flex-direction: column;
	align-items: center;
	width: max-content;
	padding: 2rem;
	border-radius: var(--rounded-3xl);
	box-shadow: var(--shadow);
`;
const Name = styled.h3`
	font-size: var(--text-4xl);
	font-weight: var(--font-bold);
	width: max-content;
	line-height: 1.1;
	padding: 1rem;
	margin-bottom: 0.5rem;
	background-color: var(--color-gray-100);
	border-radius: var(--rounded-3xl);
`;
const Price = styled.div`
	font-family: var(--font-secondary);
	font-size: var(--text-xl);
`;
const Tools = styled.div`
	display: flex;
	gap: 0.5rem;
	margin-top: 0.75rem;
`;
const Button = styled.button`
	font-family: var(--font-primary);
	font-weight: var(--font-bold);
	padding: 0.5rem;
	transition: var(--transition);
	color: var(--color-light);
	background-color: ${(props) =>
		props.background ? props.background : 'var(--color-error)'};
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
