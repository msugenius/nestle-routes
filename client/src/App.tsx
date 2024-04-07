import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'
import Navigator from './components/Navigator'

function App() {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <Navigator />
    </QueryClientProvider>
  )
}

export default App
