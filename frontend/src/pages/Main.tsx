import { useEffect, useState } from 'react';
import axios from 'axios';
import { Pet } from '@/types/Pet';
import PetCard from '@/Components/PetCard';

const Main = () => {
	const [data, setData] = useState([]);

	useEffect(() => {
		async function fetchData() {
			const response = await axios.get('http://localhost:8080/pets');

			setData(response.data);
		}

		fetchData();
	}, []);
	return (
		<div>
			{data.map((pet: Pet) => (
				<PetCard key={pet.Id} pet={pet} />
			))}
		</div>
	);
};

export default Main;
