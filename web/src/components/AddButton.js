import React from 'react';
import styled from 'styled-components';
import AddIcon from '../icons/AddIcon';

const AddButton = ({ onClick, height = 1.5, width = 1.5 }) => {
	return (
		<Button onClick={onClick}>
			<AddIcon height={height} width={width} />
		</Button>
	);
};

const Button = styled.button`
	display: flex;
	padding: 2rem;
	justify-content: center;
	align-items: center;

	cursor: pointer;

	color: var(--color-gray-500);
	background-color: var(--color-gray-100);
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

export default AddButton;
