import styled from 'styled-components';

const Location = ({ value = {} }) => {
	return (
		<LocationContainer>
			<Info>
				<Label>Street</Label>
				<Text>{value.street}</Text>
			</Info>
			<Info>
				<Label>City</Label>
				<Text>{value.city}</Text>
			</Info>
			<Info>
				<Label>Province</Label>
				<Text>{value.province}</Text>
			</Info>
			<Info>
				<Label>Country</Label>
				<Text>{value.country}</Text>
			</Info>
		</LocationContainer>
	);
};

const LocationContainer = styled.div`
	display: flex;
	flex-direction: column;
	gap: 0.5rem;
`;
const Info = styled.div``;
const Label = styled.div`
	font-weight: var(--font-bold);
`;
const Text = styled.div`
	font-family: var(--font-secondary);
	font-size: var(--text-sm);
`;

export default Location;
