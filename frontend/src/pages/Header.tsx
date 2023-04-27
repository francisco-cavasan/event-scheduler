import Router from 'next/router';

const Header = () => {
	const handleLogout = () => {
		localStorage.removeItem('token');
		Router.push('/login');
	};
	return (
		<header className='flex items-center justify-between py-4 px-6 bg-white'>
			<div className='flex items-center'>
				<img
					src='https://static.vecteezy.com/ti/vetor-gratis/p3/6470722-vector-pet-shop-logo-design-template-modern-animal-icon-label-for-store-veterinary-clinic-hospital-shelter-business-services-flat-illustration-background-with-dog-cat-and-cavalo-gratis-vetor.jpg'
					alt='logo'
					className='h-16 w-16 mr-2'
				/>
			</div>
			<div className='flex items-center'>
				<button className='bg-blue-500 text-white py-2 px-4 rounded mr-2'>Meus Pets</button>
				<button className='bg-blue-500 text-white py-2 px-4 rounded mr-2'>Perfil</button>
				<button className='bg-red-400 text-white py-2 px-4 rounded mr-2' onClick={handleLogout}>
					Sair
				</button>
			</div>
		</header>
	);
};

export default Header;
