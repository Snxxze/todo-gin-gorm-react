import api from "./api"
import type { LoginRequest, RegisterRequest } from "../interfaces/Auth"

// register
export const register = async (data: RegisterRequest) => {
  const res = await api.post("/auth/register", data)
  return res.data;
}

// login
export const login = async (data: LoginRequest) => {
  const res = await api.post("/auth/login", data)
  return res.data
}

// logout
export const logout = () => {
  localStorage.removeItem("token")
}