import React, { useState } from 'react';
import NumberFormat from 'react-number-format';
import Header from '../elements/Header';
import Input from '../elements/Input';
import InputContainer from '../elements/InputContainer';
import Label from '../elements/Label';
import SubmitButton from '../elements/SubmitButton';
import useAddInventory from '../hooks/useAddInventory';
import useUpdateInventory from '../hooks/useUpdateInventory';

const InventoryForm = ({ value, onSubmit = () => {} }) => {
	const [name, setName] = useState((value && value.name) || '');
	const [price, setPrice] = useState((value && value.price) || 0);
	const [isValid, setIsValid] = useState(false);

	const { mutate: addInventory } = useAddInventory();
	const { mutate: updateInventory } = useUpdateInventory();

	const validate = (name, price) => {
		if (name !== '' && price >= 0) {
			setIsValid(true);
		} else {
			setIsValid(false);
		}
	};

	const handleNameChange = (event) => {
		validate(event.target.value, price);
		setName(event.target.value);
	};

	const handlePriceChange = (values) => {
		const { value } = values;
		const input = parseFloat(value);
		validate(name, input);
		setPrice(input);
	};
	const handleSubmit = () => {
		const inventory = {
			name,
			price,
		};
		if (value) updateInventory({ id: value._id, inventory });
		else addInventory(inventory);
		onSubmit();
	};

	return (
		<div>
			<Header>{value ? 'Update Inventory' : 'Add Inventory'}</Header>
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
			<InputContainer>
				<Label>
					Price
					<NumberFormat
						value={price}
						onValueChange={handlePriceChange}
						className='foo'
						displayType='input'
						type='text'
						thousandSeparator={true}
						prefix={'$'}
						decimalScale={2}
						customInput={Input}
					/>
				</Label>
			</InputContainer>
			<SubmitButton
				type='submit'
				value='Submit'
				onClick={handleSubmit}
				disabled={!isValid}
			>
				{value ? 'Update' : 'Submit'}
			</SubmitButton>
		</div>
	);
};

export default InventoryForm;
