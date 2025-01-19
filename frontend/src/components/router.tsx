import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import SubmitPage from "../page/submit";
import React from "react";
import LoginPage from "../page/login";
import TaskPage from "../page/tasks";

export default function AppRouter() {
    return <BrowserRouter>
        <Routes>
            <Route index element={<Navigate to={"/login"} />} />
            <Route path="/submit" element={<SubmitPage />} />
            <Route path="/tasks" element={<TaskPage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="*" element={<Navigate to={"/"} />} />
        </Routes>
    </BrowserRouter>
}