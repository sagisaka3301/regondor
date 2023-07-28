import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { Task } from '../types'
import { useError } from './useError'

export const useQueryTasks = () => {
  // useErrorからswitchErrorHandlingの関数を読み込む.
  const { switchErrorHandling } = useError()
  // タスク一覧を取得するために、getTaskという関数を定義。axiosのgetメソッドでtasksのエンドポイントにアクセスしタスク一覧を取得できるようにしておく。
  const getTasks = async () => {
    const { data } = await axios.get<Task[]>(
      `${process.env.REACT_APP_API_URL}/tasks`,
      { withCredentials: true }
    )
    return data
  }

  // カスタムフックのreturnの値としてuseQueryを実行した結果を渡す。
  // reactQueryでは、fetchで取得したデータをクライアントのキャッシュに格納してくれる。
  return useQuery<Task[], Error>({
    queryKey: ['tasks'], // キャッシュのキーとして、tasksという名前を付ける。
    queryFn: getTasks, // 上で定義しているgetTasks関数を指定。
    staleTime: Infinity, // キャッシュしたデータをどれだけ最新のものとしてみなすか。今回はマニュアルでキャッシュを更新するので、infinity(永久)に設定。
    onError: (err: any) => {
      // エラー発生の場合はエラーメッセージを取得し、switchErrorHandling関数を呼び出す。
      if (err.response.data.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response.data)
      }
    },
  })
}
