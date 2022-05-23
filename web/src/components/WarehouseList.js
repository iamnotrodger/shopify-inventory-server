import { Link } from 'react-router-dom';
import styled from 'styled-components';
import Warehouse from './Warehouse';

const WarehouseList = ({ items = [], children }) => {
	return (
		<List>
			{items.map((warehouse, i) => (
				<Link key={i} to={`warehouse/${warehouse._id}`}>
					<Warehouse value={warehouse} />
				</Link>
			))}
			{children}
		</List>
	);
};

const List = styled.div`
	display: grid;
	grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
	gap: 0.75rem;
`;

export default WarehouseList;
