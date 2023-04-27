import React, { useEffect, useState } from 'react';

import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import { Button, FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput } from '@mui/material';
import { VisibilityOff, Visibility } from '@mui/icons-material';
import axios from 'axios';
import Router from 'next/router';

const LoginPage = () => {
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');
	const [showPassword, setShowPassword] = useState(false);
	let isLogged = false;
	if (typeof window !== 'undefined') {
		// Perform localStorage action
		isLogged = !!localStorage.getItem('token');
	}

	if (isLogged) {
		useEffect(() => {
			Router.push('/login');
		})
	}

	const handleClickShowPassword = () => {
		setShowPassword(!showPassword);
	};

	const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
		event.preventDefault();
	};

	const handleSubmit = async (e: any) => {
		e.preventDefault();
		const { data } = await axios.post('http://localhost:8080/login', { email, password });

		if (!!data.token) {
			localStorage.setItem('token', data.token);
			window.location.replace('/');
		} else {
			alert('Email ou senha incorretos');
		}
	};

	return (
		<Box
			component='form'
			sx={{
				'& .MuiTextField-root': { m: 1, width: '25ch' },
			}}
			autoComplete='off'
			display='flex'
			flexDirection='column'
			justifyContent='center'
			alignItems='center'
			minHeight='100vh'
			onSubmit={handleSubmit}
		>
			<TextField
				required
				autoComplete='email'
				id='outlined-required'
				label='Email'
				type='email'
				value={email}
				onChange={(e) => setEmail(e.target.value)}
			/>

			<FormControl sx={{ m: 1, width: '25ch' }} variant='outlined'>
				<InputLabel htmlFor='outlined-adornment-password'>Senha</InputLabel>
				<OutlinedInput
					id='outlined-adornment-password'
					type={showPassword ? 'text' : 'password'}
					value={password}
					onChange={(e) => setPassword(e.target.value)}
					endAdornment={
						<InputAdornment position='end'>
							<IconButton
								aria-label='toggle password visibility'
								onClick={handleClickShowPassword}
								onMouseDown={handleMouseDownPassword}
								edge='end'
							>
								{showPassword ? <VisibilityOff /> : <Visibility />}
							</IconButton>
						</InputAdornment>
					}
					label='Confirme a Senha'
				/>
			</FormControl>
			<Button variant='outlined' type='submit'>
				Entrar
			</Button>
		</Box>
	);
};

export default LoginPage;
