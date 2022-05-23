import React, { useState } from 'react';
import Header from '../elements/Header';
import Input from '../elements/Input';
import InputContainer from '../elements/InputContainer';
import Label from '../elements/Label';
import SubmitButton from '../elements/SubmitButton';
import useAddWarehouse from '../hooks/useAddWarehouse';

const WarehouseForm = ({ onSubmit = () => {} }) => {
	const [name, setName] = useState('');
	const [isValid, setIsValid] = useState(false);

	const { mutate: addWarehouse } = useAddWarehouse();

	const validate = (name) => {
		if (name !== '') {
			setIsValid(true);
		} else {
			setIsValid(false);
		}
	};

	const handleNameChange = (event) => {
		validate(event.target.value);
		setName(event.target.value);
	};

	const handleSubmit = () => {
		const warehouse = { name };
		addWarehouse(warehouse);
		onSubmit();
	};

	return (
		<div>
			<Header>Add Warehouse</Header>
			<InputContainer>
				<Label>
					Name
					<Input
						type='text'
						value={name}
						onChange={handleNameChange}
					/>
				</Label>
			</InputContainer>
			<SubmitButton
				type='submit'
				value='Submit'
				onClick={handleSubmit}
				disabled={!isValid}
			>
				Submit
			</SubmitButton>
		</div>
	);
};

export default WarehouseForm;
