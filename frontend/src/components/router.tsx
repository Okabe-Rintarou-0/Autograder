import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import SubmitPage from "../page/submit";
import LoginPage from "../page/login";
import TaskPage from "../page/tasks";
import UnauthorizedPage from "../page/403";
import NotfoundPage from "../page/404";
import UsersPage from "../page/users";

export default function AppRouter() {
    return <BrowserRouter>
        <Routes>
            <Route index element={<Navigate to={"/submit"} />} />
            <Route path="/submit" element={<SubmitPage />} />
            <Route path="/tasks" element={<TaskPage />} />
            <Route path="/users" element={<UsersPage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/403" element={<UnauthorizedPage />} />
            <Route path="/404" element={<NotfoundPage />} />
            <Route path="*" element={<Navigate to={"/404"} />} />
        </Routes>
    </BrowserRouter>
}