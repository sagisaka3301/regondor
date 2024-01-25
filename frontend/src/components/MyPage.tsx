import { useQueryMypage } from '../hooks/useQueryMypage'
import { useQueryClient } from '@tanstack/react-query'
import AppLayout from './AppLayout'
export const MyPage = () => {
  const queryClient = useQueryClient()

  // データがundefinedまたはloadingの場合、isLoadingがtrueになる。データ取得後にfalseになる。
  const { data, isLoading } = useQueryMypage()

  console.log(data)

  return (
    <AppLayout>
      <div className="wrap">
        <div className="profile">
          {isLoading ? (
            <p>Loading...</p>
          ) : (
            <ul>
              <p>{data ? data.email : 'No data acailable'}</p>
            </ul>
          )}
        </div>
        <p></p>
      </div>
    </AppLayout>
  )
}
