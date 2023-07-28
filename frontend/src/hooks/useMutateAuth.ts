import axios from 'axios'
import { useNavigate } from 'react-router-dom'
import { useMutation } from '@tanstack/react-query'
import useStore from '../store'
import { Credential } from '../types'
import { useError } from '../hooks/useError'

export const useMutateAuth = () => {
  const navigate = useNavigate()
  const resetEditedTask = useStore((state) => state.resetEditedTask)
  // useErrorのカスタムフックからswitchErrorHandling関数を読み込む。
  const { switchErrorHandling } = useError()

  // ログインを行うためのMutationを作る。react-queryのuseMutationを使い実装。
  const loginMutation = useMutation(
    // 引数でクレデンシャル情報(EmailとPasswordの情報)を受け取り、axiosのPOSTメソッドでloginのエンドポイントにアクセス。
    async (user: Credential) =>
      // 第二引数でクレデンシャルオブジェクトのEmailとpasswordを渡す(user)。
      await axios.post(`${process.env.REACT_APP_API_URL}/login`, user),
    {
      // ログイン成功時、todoのページに遷移。
      onSuccess: () => {
        navigate('/todo')
      },
      // エラーが発生S田場合、エラーメッセージを取り出し、switchErrorHandling関数を呼び出す。
      // csrfミドルウェア関係のエラーだけは、エラーメッセージがresponse.data.messageの階層に格納されるので、そのエラーが存在する場合はresponse.data.messageから取り出して関数を呼ぶ。
      // それ以外の場合は、response.dataからメッセージを取り出す。
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )

  // singUpするためのregisterMutationを追加。
  const registerMutation = useMutation(
    // 引数でemailとpasswordの情報を受け取り、axiosのPOSTメソッドでsignupのエンドポイントにアクセス。
    async (user: Credential) =>
      await axios.post(`${process.env.REACT_APP_API_URL}/signup`, user),
    {
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )

  // ログアウトMutation
  const logoutMutation = useMutation(
    async () => await axios.post(`${process.env.REACT_APP_API_URL}/logout`),
    {
      onSuccess: () => {
        resetEditedTask()
        navigate('/')
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )
  // カスタムフックの最後に3つ定義した関数を返すようにする。
  return { loginMutation, registerMutation, logoutMutation }
}
