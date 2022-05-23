import { Link } from 'react-router-dom';
import styled from 'styled-components';
import Inventory from './Inventory';

const WarehouseInventoryList = ({ items = [], icon, onClick, children }) => {
	return (
		<List>
			{items.map((inventory, i) => (
				<Container key={i}>
					<div onClick={() => onClick(inventory)}>{icon}</div>
					<Link to={`/inventory/${inventory._id}`}>
						<Inventory value={inventory} />
					</Link>
				</Container>
			))}
			{children}
		</List>
	);
};

const List = styled.div`
	display: grid;
	grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
	gap: 0.75rem;
	box-shadow: var(--shadow);
	border-radius: var(--rounded-3xl);
	padding: 1rem;
`;

const Container = styled.div`
	position: relative;
`;

export default WarehouseInventoryList;
