import React from 'react';
import Modal from 'react-modal';
import styled from 'styled-components';
import CloseIcon from '../icons/CloseIcon';

const customStyles = {
	content: {
		top: '50%',
		left: '50%',
		right: 'auto',
		bottom: 'auto',
		marginRight: '-50%',
		transform: 'translate(-50%, -50%)',
		padding: '0',
		borderWidth: '0',
		borderRadius: 'var(--rounded-lg)',
		boxShadow: `var(--shadow-lg)`,
	},
};

const ModalContainer = ({ onClose, children, ...props }) => {
	return (
		<Modal style={customStyles} onRequestClose={onClose} {...props}>
			<Button onClick={onClose}>
				<CloseIcon height={1.5} width={1.5} />
			</Button>
			<Container>{children}</Container>
		</Modal>
	);
};

const Container = styled.div`
	position: relative;
	padding: 3rem;
`;

const Button = styled.button`
	position: absolute;
	top: 0;
	right: 0;
	margin: 1rem;
	z-index: 50;
	background: none;

	cursor: pointer;
`;

export default ModalContainer;
