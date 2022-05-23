import { Link } from 'react-router-dom';
import styled from 'styled-components';
import Inventory from './Inventory';

const InventoryList = ({ items = [], children }) => {
	return (
		<List>
			{items.map((inventory, i) => (
				<Link key={i} to={`/inventory/${inventory._id}`}>
					<Inventory value={inventory} />
				</Link>
			))}
			{children}
		</List>
	);
};

const List = styled.div`
	display: grid;
	grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
	gap: 0.75rem;
`;

export default InventoryList;
