import Modal from 'react-modal';
import { QueryClient, QueryClientProvider } from 'react-query';
import { Route, Routes } from 'react-router-dom';
import HomePage from './pages/HomePage';

const queryClient = new QueryClient();
Modal.setAppElement('#root');

const App = () => {
	return (
		<QueryClientProvider client={queryClient}>
			<Routes>
				<Route index element={<HomePage />} />
			</Routes>
		</QueryClientProvider>
	);
};

export default App;
