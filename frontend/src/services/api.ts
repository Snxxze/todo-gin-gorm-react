import axios from "axios";

// กำหนด api backend ไว้ใช้เรียก Api 
const api = axios.create({
  baseURL: "http://localhost:8000",
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config;
})

export default api;