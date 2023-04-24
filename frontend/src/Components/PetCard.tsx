import React from 'react';
import {
	Card,
	CardMedia,
	CardContent,
	Typography,
	List,
	ListItem,
	ListItemText,
	Button,
} from '@mui/material';
import { Pet } from '@/types/Pet';
import axios from 'axios';

const onPetFoundHandler = async (petId) => {
	await axios.post(`http://localhost:8080/pets/found`, {
    id: petId
  });
};

const PetCard = (petProp: Pet) => {
	const pet = petProp.pet;

	return (
		<Card sx={{ width: 380 }}>
			<CardMedia
				component='img'
				sx={{
					height: 180,
				}}
				image={pet.Image_url}
				alt={pet.Name}
			/>
			<CardContent>
				<Typography gutterBottom variant='h5' component='div'>
					<strong>{pet.Name}</strong>
				</Typography>
				<List>
					<ListItem>
						<ListItemText primary={`Idade: ${pet.Age}`} />
					</ListItem>
					<ListItem>
						<ListItemText primary={`Descrição: ${pet.Description}`} />
					</ListItem>
					<ListItem>
						<ListItemText primary={`Localização: ${pet.Location}`} />
					</ListItem>
					<ListItem>
						<Button variant='outlined' onClick={(e) => {onPetFoundHandler(pet.Id)}}>
							Marcar como encontrado
						</Button>
					</ListItem>
				</List>
			</CardContent>
		</Card>
	);
};

export default PetCard;
