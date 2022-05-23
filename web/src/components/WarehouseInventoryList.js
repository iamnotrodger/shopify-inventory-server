import React from 'react';
import styled from 'styled-components';
import InventoryList from './InventoryList';

const WarehouseInventoryList = ({ items = [] }) => {
	const handleClick = () => {};

	return (
		<List>
			{items.map((inventory, i) => (
				<InventoryList key={i} value={inventory} />
			))}
		</List>
	);
};

const List = styled.div`
	display: grid;
	grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
	gap: 0.75rem;
`;

export default WarehouseInventoryList;
