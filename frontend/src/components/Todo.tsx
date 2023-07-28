import { FormEvent } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import {
  ArrowRightOnRectangleIcon,
  ShieldCheckIcon,
} from '@heroicons/react/24/solid'
import useStore from '../store'
import { useQueryTasks } from '../hooks/useQueryTasks'
import { useMutateTask } from '../hooks/useMutateTask'
import { useMutateAuth } from '../hooks/useMutateAuth'
import { TaskItem } from './TaskItem' // taskItemコンポーネントもimportする。

export const Todo = () => {
  // useQueryClientのzastandからeditTaskのstateとupdatedTaskの関数を読み込み。
  const queryClient = useQueryClient()
  const { editedTask } = useStore()
  const updateTask = useStore((state) => state.updateEditedTask)
  // useQueryTasksのカスタムフックからdataとisLoadingを読み込む。
  const { data, isLoading } = useQueryTasks()
  // useMutateTaskのカスタムフックから以下2つの関数を読み込み。
  const { createTaskMutation, updateTaskMutation } = useMutateTask()
  // logoutMutationの読み込み。
  const { logoutMutation } = useMutateAuth()
  // submitボタンが押されたときに実行される関数
  const submitTaskHandler = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    // zastandステートのeditedTaskのidが0の場合、createとみなして、createTaskMutationを実行
    if (editedTask.id === 0)
      createTaskMutation.mutate({
        title: editedTask.title, // このときに渡す値は、zastandのeditedTaskステートに格納されているタイトルを渡す。
      })
    else {
      // 0以外の場合は、アップデートとみなす。zastandのステートであるeditedTaskステートを引数で渡す。
      updateTaskMutation.mutate(editedTask)
    }
  }

  // logout関数を定義し、関数内でlogoutMutationとmutateAsyncを呼び出す。
  const logout = async () => {
    await logoutMutation.mutateAsync()
    queryClient.removeQueries(['tasks']) // アプリケーションをログアウトするときに、クライアントサイドの['tasks']のキーワードで格納されているキャッシュをクリアする必要がある。tasksのキーワードに紐づいたキャッシュを削除する。
  }
  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
      <div className="flex items-center my-3">
        <ShieldCheckIcon className="h-8 w-8 mr-3 text-indigo-500 cursor-pointer" />
        <span className="text-center text-3xl font-extrabold">TaskManager</span>
      </div>
      <ArrowRightOnRectangleIcon
        onClick={logout}
        className="h-6 w-6 my-6 text-blue-500 cursor-pointer"
      />
      {/* ユーザーがタスクのタイトルを入力するためのフォーム */}
      <form onSubmit={submitTaskHandler}>
        <input
          className="mb-3 mr-3 px-3 py-2 border border-gray-300"
          placeholder="title ?"
          type="text"
          onChange={(e) => updateTask({ ...editedTask, title: e.target.value })}
          value={editedTask.title || ''}
        />
        <button
          className="disabled:opacity-40 mx-3 py-2 px-3 text-white bg-indigo-600 rounded"
          disabled={!editedTask.title}
        >
          {editedTask.id === 0 ? 'Create' : 'Update'}
        </button>
      </form>
      {/* タスクの一覧を表示 */}
      {/* useQueryのisLoadingステートがtrueの時はLoadingと表示。fetchが終わったら取得したデータをマップで展開し、taskItemコンポーネントにpropsとしてタスクのidとタスクのtitleを渡す。 */}
      {isLoading ? (
        <p>Loading...</p>
      ) : (
        <ul className="my-5">
          {data?.map((task) => (
            <TaskItem key={task.id} id={task.id} title={task.title} />
          ))}
        </ul>
      )}
    </div>
  )
}
