import React from 'react';
import styled from 'styled-components';
import WarehouseIcon from '../icons/WarehouseIcon';

const Warehouse = ({ value = {} }) => {
	return (
		<Container>
			<IconContainer>
				<WarehouseIcon height={1.5} width={1.5} />
			</IconContainer>
			<Info>
				<Name>{value.name}</Name>
				<Location>{value.location && value.location.street}</Location>
			</Info>
		</Container>
	);
};

const Container = styled.div`
	display: flex;
	align-items: center;
	gap: 0.75rem;
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
const Location = styled.div`
	font-family: var(--font-secondary);
	font-size: var(--text-base);
`;

export default Warehouse;
