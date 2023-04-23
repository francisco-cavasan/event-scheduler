import React from "react";
import {
  Card,
  CardMedia,
  CardContent,
  Typography,
  List,
  ListItem,
  ListItemText,
} from "@mui/material";
import { Pet } from "@/types/Pet";

const PetCard = (pet: Pet) => {
  return (
    <Card sx={{ maxWidth: 345 }}>
      <CardMedia component="img" height="140" image={pet.ImageUrl} alt={pet.Name} />
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          <strong>{pet.Name}</strong>
        </Typography>
        <List>
          <ListItem>
            <ListItemText primary={`Age: ${pet.Age}`} />
          </ListItem>
          <ListItem>
            <ListItemText primary={`Description: ${pet.Description}`} />
          </ListItem>
          <ListItem>
            <ListItemText primary={`Location: ${pet.Location}`} />
          </ListItem>
        </List>
      </CardContent>
    </Card>
  );
};

export default PetCard;
