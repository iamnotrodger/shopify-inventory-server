import styled from 'styled-components';
import InventoryIcon from '../icons/InventoryIcon';

const Inventory = ({ value = {} }) => {
	return (
		<Container>
			<IconContainer>
				<InventoryIcon height={1.5} width={1.5} />
			</IconContainer>
			<Info>
				<Name>{value.name}</Name>
				<Price>${value.price}</Price>
			</Info>
		</Container>
	);
};

const Container = styled.div`
	display: flex;
	gap: 0.5rem;
	padding: 2rem;
	border-radius: var(--rounded-xl);
	box-shadow: var(--shadow);
	transition: var(--transition);

	:hover {
		box-shadow: var(--shadow-lg);
	}
	:active {
		box-shadow: var(--shadow);
	}
`;
const IconContainer = styled.div`
	align-self: center;
`;
const Info = styled.div``;
const Name = styled.div`
	font-family: var(--font-primary);
	font-size: var(--text-lg);
	font-weight: var(--font-bold);
	line-height: 1;
`;
const Price = styled.div`
	font-family: var(--font-secondary);
	font-size: var(--text-sm);
`;

export default Inventory;
