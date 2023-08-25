import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { MyPage } from '../types'
import { useError } from './useError'

export const useQueryMypage = () => {
  const { switchErrorHandling } = useError()

  // ユーザー情報を取得しマイページに渡す。
  // ※1件(単一のオブジェクトなので、<MyPage[]>ではなく、<MyPage>)
  const getUser = async () => {
    const { data } = await axios.get<MyPage>(
      `${process.env.REACT_APP_API_URL}/mypage`,
      { withCredentials: true }
    )
    return data
  }

  return useQuery<MyPage, Error>({
    queryKey: ['user'],
    queryFn: getUser,
    staleTime: Infinity,
    onError: (err: any) => {
      if (err.response.data.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response.data)
      }
    },
  })
}
