import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import SubmitPage from "../page/submit";
import React from "react";

export default function AppRouter() {
    return <BrowserRouter>
        <Routes>
            <Route index element={<Navigate to={"/submit"} />} />
            <Route path="/submit" element={<SubmitPage />} />
            <Route path="*" element={<Navigate to={"/"} />} />
        </Routes>
    </BrowserRouter>
}