import {
  BrowserRouter,
  Routes,
  Route
} from "react-router-dom"
import LoginPage from "../pages/LoginPage"
import RegisterPage from "../pages/RegisterPage"
import TodoPage from "../pages/TodoPage"

const AppRoutes = () => (
  <BrowserRouter>
    <Routes>
      <Route path="/login" element={<LoginPage/>}/>
      <Route path="/register" element={<RegisterPage/>}/>
      <Route path="/todos" element={<TodoPage />} />
      <Route path="*" element={<LoginPage/>}/>
    </Routes>
  </BrowserRouter>
)

export default AppRoutes