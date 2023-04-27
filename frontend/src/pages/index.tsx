import Header from './Header';
import { Box } from '@mui/material';
import Main from './Main';
import { useEffect } from 'react';
import Router from 'next/router';

export default function Home() {
	let isLogged = false;
	if (typeof window !== 'undefined') {
		// Perform localStorage action
		isLogged = !!localStorage.getItem('token');
	}

	if (!isLogged) {
		useEffect(() => {
			Router.push('/login');
		})
	}

	return (
		<Box
			sx={{
				width: '100%',
				backgroundColor: '#ebeae6',
			}}
		>
			<Header />
			<Main />
		</Box>
	);
}
