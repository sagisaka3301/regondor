import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App'
import reportWebVitals from './reportWebVitals'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'

// Reactクエリーのクエリークライアントをnew QueryClientで作成する。
const queryClient = new QueryClient({})

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
root.render(
  <React.StrictMode>
    {/* アプリ全体にreactQueryClientを適用するために、AppコンポーネントをQueryClientProviderでラップする。 */}
    <QueryClientProvider client={queryClient}>
      <App />
      {/* Reactクエリーのdevelopmenttoolsを有効化するために、以下の一行を追加。 */}
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  </React.StrictMode>
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
