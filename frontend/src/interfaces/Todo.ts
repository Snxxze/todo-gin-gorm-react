export interface Todo {
  id: number;
  title: string;
  description?: string;
  status: "pending" | "done";
  createdAt: string;
  updatedAt: string;
  userId: number;
}

export interface CreateTodoRequest {
  title: string;
  description?: string;
}

export interface UpdateTodoRequest {
  title?: string;
  description?: string;
  status?: "pending" | "done";
}