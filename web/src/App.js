import Modal from 'react-modal';
import { QueryClient, QueryClientProvider } from 'react-query';
import { Route, Routes } from 'react-router-dom';
import Title from './elements/Title';
import HomePage from './pages/HomePage';
import InventoryPage from './pages/InventoryPage';
import WarehousePage from './pages/WarehousePage';

const queryClient = new QueryClient();
Modal.setAppElement('#root');

const App = () => {
	return (
		<QueryClientProvider client={queryClient}>
			<Title>Shopify-Inventory</Title>
			<Routes>
				<Route index element={<HomePage />} />
				<Route path='inventory/:id' element={<InventoryPage />} />
				<Route path='warehouse/:id' element={<WarehousePage />} />
			</Routes>
		</QueryClientProvider>
	);
};

export default App;
