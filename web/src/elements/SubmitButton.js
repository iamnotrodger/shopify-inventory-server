import styled from 'styled-components';

export default styled.button`
	font-family: var(--font-primary);
	font-size: var(--text-lg);
	font-weight: var(--font-bold);
	width: 100%;
	padding: 1rem;
	margin-top: 0.5rem;
	border-radius: var(--rounded-full);
	transition: var(--transition);

	cursor: pointer;

	:enabled {
		color: var(--color-light);
		background-color: var(--color-success);
		box-shadow: var(--shadow);
	}

	:disabled {
		cursor: not-allowed;
	}

	:hover:enabled {
		box-shadow: var(--shadow-lg);
	}

	:active:enabled {
		box-shadow: var(--shadow);
	}
`;
