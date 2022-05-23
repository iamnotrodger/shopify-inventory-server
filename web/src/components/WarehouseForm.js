import React, { useState } from 'react';
import Header from '../elements/Header';
import Input from '../elements/Input';
import InputContainer from '../elements/InputContainer';
import Label from '../elements/Label';
import SubmitButton from '../elements/SubmitButton';
import useAddWarehouse from '../hooks/useAddWarehouse';

const WarehouseForm = ({ onSubmit = () => {} }) => {
	const [name, setName] = useState('');
	const [street, setStreet] = useState('');
	const [city, setCity] = useState('');
	const [province, setProvince] = useState('');
	const [country, setCountry] = useState('');
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

	const handleStreetChange = (event) => {
		setStreet(event.target.value);
	};

	const handleCityChange = (event) => {
		setCity(event.target.value);
	};

	const handleProvinceChange = (event) => {
		setProvince(event.target.value);
	};

	const handleCountryChange = (event) => {
		setCountry(event.target.value);
	};

	const handleSubmit = () => {
		const warehouse = {
			name,
			location: {
				street,
				city,
				province,
				country,
			},
		};
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
			<InputContainer>
				<Label>
					Street
					<Input
						type='text'
						value={street}
						onChange={handleStreetChange}
					/>
				</Label>
			</InputContainer>
			<InputContainer>
				<Label>
					City
					<Input
						type='text'
						value={city}
						onChange={handleCityChange}
					/>
				</Label>
			</InputContainer>
			<InputContainer>
				<Label>
					Province
					<Input
						type='text'
						value={province}
						onChange={handleProvinceChange}
					/>
				</Label>
			</InputContainer>
			<InputContainer>
				<Label>
					Country
					<Input
						type='text'
						value={country}
						onChange={handleCountryChange}
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
