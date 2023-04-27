import { useEffect, useState } from 'react';
import axios from 'axios';
import PetCard from '@/components/PetCard';
import { Pet } from '@/types/Pet';

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
		<div className='container px-6 py-3 mx-auto flex flex-wrap gap-10'>
			{data.map((pet: Pet) => (
				<PetCard key={pet.ID} pet={pet} />
			))}
		</div>
	);
};

export default Main;
