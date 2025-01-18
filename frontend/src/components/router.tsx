import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import SubmitPage from "../page/submit";

export default function AppRouter() {
    return <BrowserRouter>
        <Routes>
            <Route index element={<Navigate to={"/submit"} />} />
            <Route path="/submit" element={<SubmitPage />} />
            <Route path="*" element={<Navigate to={"/"} />} />
        </Routes>
    </BrowserRouter>
}