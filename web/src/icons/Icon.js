import styled from 'styled-components';

export default styled.svg`
	height: ${(props) => props.height || 1}rem;
	width: ${(props) => props.width || 1}rem;
	color: ${(props) => props.color || 'inherit'};
`;
