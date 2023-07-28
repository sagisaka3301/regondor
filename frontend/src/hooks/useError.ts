import axios from 'axios'
import { useNavigate } from 'react-router-dom'
import { CsrfToken } from '../types'
import useStore from '../store'

// functionalコンポーネントuseErrorを作成
export const useError = () => {
  const navigate = useNavigate()
  // zustandのstoreからresetEditTaskの関数を読み込み、コンポーネントで使用できるようにしておく。
  const resetEditedTask = useStore((state) => state.resetEditedTask)

  // 関数内でaxiosのgetメソッドでcsrfトークンのエンドポイントにアクセスする。
  // そして、取得したcsrfトークンをaxiosのdefaults.headersに設定するようにしている。
  const getCsrfToken = async () => {
    const { data } = await axios.get<CsrfToken>(
      `${process.env.REACT_APP_API_URL}/csrf`
    )
    axios.defaults.headers.common['X=CSRF-TOKEN'] = data.csrf_token
  }

  // 引数で受け取るエラーメッセージの内容に応じてswitch文を使って処理を切り替える。
  const switchErrorHandling = (msg: string) => {
    switch (msg) {
      // エラーがinvalid csrf tokenなら、getCsrfTokenを呼び出し、再度csrfトークンを取得するようにしている。
      case 'invalid csrf token':
        getCsrfToken()
        alert('CSRF toke is invalid, please try again')
        break
      // jwtがダメだったときは、zustandのステートをリセットしてindexのページにナビゲートする。
      case 'invalid or expired jwt':
        alert('access token expired, please login')
        resetEditedTask()
        navigate('/')
        break
      case 'missing or malformed jwt':
        alert('access token is not valid, please login')
        resetEditedTask()
        navigate('/')
        break
      case 'duplicated key not allowed':
        alert('email already exist, please use another one')
        break
      // password関連のエラーの場合はpassword is not correctを指定。
      case 'crypto/bcrypt: hashedPassword is not the hash of the given password':
        alert('password is not correct')
        break
      case 'record not found':
        alert('email is not correct')
        break
      default:
        alert(msg)
    }
  }
  // カスタムフックのreturnとしてswitchErrorHandlingの関数を返す。
  return { switchErrorHandling }
}
