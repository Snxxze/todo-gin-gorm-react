import api from "./api";
import type { CreateTodoRequest, Todo, UpdateTodoRequest } from "../interfaces/Todo";

// create todo
export const create = async (data: CreateTodoRequest): Promise<Todo> => {
  const  res = await api.post("/todos", data)
  return res.data.todo
}

// update todo
export const update = async (id: number, data: UpdateTodoRequest): Promise<Todo> => {
  const res = await api.put(`/todos/${id}`, data)
  return res.data.todo
}

// get
export const get = async (): Promise<Todo[]> => {
  const res = await api.get("/todos")
  return res.data.todos
}

// delete
export const deleteTodo = async (id: number): Promise<void> => {
  await api.delete(`/todos/${id}`)
}