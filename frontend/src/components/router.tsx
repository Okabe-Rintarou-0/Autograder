import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import UnauthorizedPage from "../page/403";
import NotfoundPage from "../page/404";
import LoginPage from "../page/login";
import SqlPage from "../page/sql";
import SubmitPage from "../page/submit";
import TaskPage from "../page/tasks";
import TestcasesPage from "../page/testcase";
import UsersPage from "../page/users";

export default function AppRouter() {
    return <BrowserRouter>
        <Routes>
            <Route index element={<Navigate to={"/submit"} />} />
            <Route path="/submit" element={<SubmitPage />} />
            <Route path="/tasks" element={<TaskPage />} />
            <Route path="/sql" element={<SqlPage />} />
            <Route path="/testcases" element={<TestcasesPage />} />
            <Route path="/users" element={<UsersPage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/403" element={<UnauthorizedPage />} />
            <Route path="/404" element={<NotfoundPage />} />
            <Route path="*" element={<Navigate to={"/404"} />} />
        </Routes>
    </BrowserRouter>
}