import { useState, FormEvent } from 'react'
import { CheckBadgeIcon, ArrowPathIcon } from '@heroicons/react/24/solid'
import { useMutateAuth } from '../hooks/useMutateAuth'
import GuestLayout from './GuestLayout'

export const Auth = () => {
  // useStateを使ってstateを定義する。
  const [email, setEmail] = useState('') // string
  const [name, setName] = useState('')
  const [pw, setPw] = useState('') // string
  const [isLogin, setIsLogin] = useState(true) // boolean
  const { loginMutation, registerMutation } = useMutateAuth() // useMutateAuthのカスタムフックから2つの関数を読み込み。

  // submitボタンが押されたときに実行される関数を定義
  const submitAuthHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    // isLoginがtureの場合loginMutation.mutateでloginMutationを実行。
    // emailとpasswordのステートを引数で渡す。
    if (isLogin) {
      loginMutation.mutate({
        email: email,
        password: pw,
      })
      // isLoginがfalseの場合、registerMutationとmutateAsyncでregisterMutationを呼び出す。
    } else {
      await registerMutation
        .mutateAsync({
          email: email,
          name: name,
          password: pw,
        })
        // registerに成功した場合は、続けてloginMutationを呼び出すことで、自動ログインするようにする。
        .then(() =>
          loginMutation.mutate({
            email: email,
            password: pw,
          })
        )
    }
  }
  return (
    <GuestLayout>
      <div className="flex justify-center items-center flex-col min-h-screen font-mono">
        <div className="flex items-center">
          <CheckBadgeIcon className="h-8 w-8 mr-2 text-white" />
          <span className="text-center text-3xl font-extrabold">
            Todo app by React/Go(Echo)
          </span>
        </div>
        {/* isLoginのstateの値に応じて、表示の変更 */}
        <h2 className="my-6">{isLogin ? 'Login' : 'Create a new account'}</h2>
        {/* 入力フォーム */}
        <form onSubmit={submitAuthHandler}>
          <div>
            <input
              className="mb-3 px-3 text-sm py-2 border border-gray-300"
              name="email"
              type="email"
              autoFocus
              placeholder="Email address"
              onChange={(e) => setEmail(e.target.value)}
              value={email}
            />
          </div>
          {!isLogin && (
            <div>
              <input
                className="mb-3 px-3 text-sm py-2 border border-gray-300"
                name="name"
                type="text"
                placeholder="Name"
                onChange={(e) => setName(e.target.value)}
                value={name}
              />
            </div>
          )}
          <div>
            <input
              className="mb-3 px-3 text-sm py-2 border border-gray-300"
              name="password"
              type="password"
              placeholder="Password"
              onChange={(e) => setPw(e.target.value)}
              value={pw}
            />
          </div>
          <div className="flex justify-center my-2">
            <button
              className="disabled:opacity-40 py-2 px-4 rounded text-white bg-indigo-600"
              disabled={!email || !pw} // emailかpasswordが一つでも空の場合はボタンを無効化
              type="submit"
            >
              {isLogin ? 'Login' : 'Sign Up'}
            </button>
            {/* ログインと新規登録を切り替えられるようなアイコンボタン */}
            <ArrowPathIcon
              onClick={() => setIsLogin(!isLogin)}
              className="h-6 w-6 my-2 text-blue-500 cursor-pointer"
            />
          </div>
        </form>
      </div>
    </GuestLayout>
  )
}
