import { useEffect, useState } from 'react'
// import logo from './logo.svg';
// import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Auth } from './components/Auth'
import { Todo } from './components/Todo'
import { MyPage } from './components/MyPage'
import { Entrance } from './components/Entrance'
import axios from 'axios'
import { CsrfToken } from './types'

// Appコンポーネントはアプリが起動するときに実行される。
function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true
    // 関数を定義。
    // axiosのgetメソッドを使ってcsrfのエンドポイントにアクセスする。
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf`
      )
      // 取得したcsrfトークンをaxiosのデフォルトheadersを使ってX-CSRF-Tokenというヘッダの名前を付けて付与する。
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    // 関数の実行
    getCsrfToken()
  }, [])

  const [isDark, setIsDark] = useState(false)
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Entrance />} />
        <Route path="/auth" element={<Auth />} />
        <Route path="/todo" element={<Todo />} />
        <Route path="/mypage" element={<MyPage />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
