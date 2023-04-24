import Header from './Header';
import { Box } from '@mui/material';
import Main from './Main';

export default function Home() {
	return (
		<Box sx={{
			width: '100%',
			backgroundColor: '#ebeae6',
		}}>
		<Header />
		<Main />
		</Box>
	);
}
