import Header from './Header';
import { Box } from '@mui/material';
import Main from './Main';

export default function Home() {
	return (
		<Box sx={{
			width: '100%',
			height: '100vh',
			backgroundColor: '#ebeae6',
		}}>
		<Header />
		<Main />
		</Box>
	);
}
