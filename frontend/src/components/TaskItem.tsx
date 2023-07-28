import { FC, memo } from 'react'
import { PencilIcon, TrashIcon } from '@heroicons/react/24/solid'
import useStore from '../store'
import { Task } from '../types'
import { useMutateTask } from '../hooks/useMutateTask'

// コンポーネントを作成。
// このコンポーネントは親のコンポーネントのステートの変化による再レンダリングを防ぐために、メモ化を行っている。
// コンポーネントのプロップスの型として、Taskデータ型からcreated_at,updated_atを取り除いたid, titleを受け取れるようにしておく。
const TaskItemMemo: FC<Omit<Task, 'created_at' | 'updated_at'>> = ({
  id,
  title,
}) => {
  // functionalコンポーネント内。
  // zustandのstoreからupdateEditedTask関数を読み込む。
  // useMutedTaskのカスタムフックからdeleteTaskMutation関数を読み込む。
  const updateTask = useStore((state) => state.updateEditedTask)
  const { deleteTaskMutation } = useMutateTask()

  return (
    <li className="my-3">
      {/* プロップスで受け取っていたtitleを表示。 */}
      <span className="font-bold">{title}</span>
      <div className="flex float-right ml-20">
        {/* ペンのアイコンがクリックされたとき、zastandのupdateTask関数を呼び出し、プロップスで受け取っていたidとtitleをzastandのstoreのstateにセットする。 */}
        <PencilIcon
          className="h-5 w-5 mx-1 text-blue-500 cursor-pointer"
          onClick={() => {
            updateTask({
              id: id,
              title: title,
            })
          }}
        />
        {/* ゴミ箱アイコンがクリックされたときは、deleteTaskMutationを実行(引数はpropsで受け取ったid) */}
        <TrashIcon
          className="h-5 w-5 text-blue-500 cursor-pointer"
          onClick={() => {
            deleteTaskMutation.mutate(id)
          }}
        />
      </div>
    </li>
  )
}
export const TaskItem = memo(TaskItemMemo)
