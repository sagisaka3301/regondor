export type Task = {
  id: number
  title: string
  created_at: Date
  updated_at: Date
}
export type CsrfToken = {
  csrf_token: string
}
export type Credential = {
  email: string
  name: string
  password: string
}

export type MyPage = {
  id: number
  email: string
  name: string
  updated_at: Date
}

export type Login = {
  email: string
  password: string
}
