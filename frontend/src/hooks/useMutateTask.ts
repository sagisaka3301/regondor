import axios from 'axios'
import { useQueryClient, useMutation } from '@tanstack/react-query'
import { Task } from '../types'
import useStore from '../store'
import { useError } from './useError'

export const useMutateTask = () => {
  // userQueryClientを変数に格納。
  const queryClient = useQueryClient()
  const { switchErrorHandling } = useError() // useErrorのカスタムフックからswitchErrorHandlingの関数を読み込み。
  const resetEditedTask = useStore((state) => state.resetEditedTask) // zustandのストアからresetEditedTask関数を読み込み、変数に格納。

  // タスクを新規作成するためのmutationを作成。
  // react-queryのuseMutationを使う。
  // 引数のデータ型として、Taskのデータ型からid, created_at, updated_atを取り除いたTaskの型を指定。
  // 引数で受け取ったTaskのオブジェクトをaxiosのPOSTメソッドでtasksのエンドポイントにリクエストを投げる。
  const createTaskMutation = useMutation(
    (task: Omit<Task, 'id' | 'created_at' | 'updated_at'>) =>
      axios.post<Task>(`${process.env.REACT_APP_API_URL}/tasks`, task),
    {
      // createに成功した場合は、getQueryDataでキャッシュの中に格納されているタスクの一覧を['tasks']のキーワード使って取得。
      onSuccess: (res) => {
        const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
        // 既存のキャッシュが存在する場合、その配列の末尾に新しく作成したタスクの要素を追加し、それを新しい配列としてキャッシュの内容を更新する。
        if (previousTasks) {
          queryClient.setQueryData(['tasks'], [...previousTasks, res.data])
        }
        // resetEditedTask関数を呼び出して、zustandのstateをリセットする。
        resetEditedTask()
      },
      // 失敗した場合は、switchErrorHandling関数を呼び出す。
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )

  // アップデート用のmutation
  const updateTaskMutation = useMutation(
    // 引数の型：Taskからcreated_atとupdated_atを除いたidとtitleの型を設定。
    (task: Omit<Task, 'created_at' | 'updated_at'>) =>
      // axiosのPUTメソッドを使い、パス末尾にタスクのidを付与。そのデータとしてタスクのタイトルを渡す。
      axios.put<Task>(`${process.env.REACT_APP_API_URL}/tasks/${task.id}`, {
        title: task.title,
      }),
    {
      onSuccess: (res, variables) => {
        // 成功した場合、既存のキャッシュの内容を取得。
        // 既存のキャッシュが存在する場合は、その配列をmapで展開する。
        const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
        if (previousTasks) {
          // 更新したタスクのidに一致するオブジェクトの要素を更新後のtaskに書き換える。
          // 新しく配列を作り、その内容でキャッシュを更新する。
          queryClient.setQueryData<Task[]>(
            ['tasks'],
            previousTasks.map((task) =>
              task.id === variables.id ? res.data : task
            )
          )
        }
        resetEditedTask()
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

  // deleteTask
  // 引数として削除したいtaskのidを受け取る。
  // axiosのdeleteメソッドでtasksの末尾にidを指定する。
  const deleteTaskMutation = useMutation(
    (id: number) =>
      axios.delete(`${process.env.REACT_APP_API_URL}/tasks/${id}`),
    {
      // 成功した場合：既存のキャッシュの配列に対し、フィルターをかけ、今削除したオブジェクトだけを取り除く。
      onSuccess: (_, variables) => {
        const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
        if (previousTasks) {
          queryClient.setQueryData<Task[]>(
            ['tasks'],
            previousTasks.filter((task) => task.id !== variables)
          )
        }
        resetEditedTask()
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

  // 定義した3つの関数を実行する。
  return {
    createTaskMutation,
    updateTaskMutation,
    deleteTaskMutation,
  }
}
