import { create } from 'zustand'

// 管理したいステートをEditedTaskという名前で定義しておく。
type EditedTask = {
  id: number
  title: string
}

// Stateを作成する。
type State = {
  // 状態管理したいステートをeditedTaskという名前で定義し、データ型をEditedTaskに指定。
  editedTask: EditedTask
  // ステートを更新するためのupdateEditedTask関数の型を定義：EditedTask型のペイロードを引数で受け取り、返り値はvoidにしている。
  updateEditedTask: (payload: EditedTask) => void
  // ステートをリセットするためのresetEditedTaskの型を定義。：引数なしで返り値void
  resetEditedTask: () => void
}

// zustandのcreateを使って、ステートと関数の具体的な処理を実装。

const useStore = create<State>((set) => ({
  editedTask: { id: 0, title: '' },
  // updateEditedTaskの具体的な処理。受け取ったpayloadをsetという関数を使ってEditedTaskのステートに設定する。
  updateEditedTask: (payload) =>
    set({
      editedTask: payload,
    }),
  // setを使ってEditedTaskのステートを初期値に初期化できるようにしておく。
  resetEditedTask: () => set({ editedTask: { id: 0, title: '' } }),
}))

export default useStore
