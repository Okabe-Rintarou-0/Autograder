import { Navigate, useNavigate } from 'react-router-dom';
import axios from "axios";
import { LoginRequest, LoginResponse, User } from "../model/user";
import { removeToken } from './token';

export async function login(request: LoginRequest) {
    const formData = new FormData();
    formData.append('identifier', request.identifier);
    formData.append('password', request.password);
    let resp = await axios.post<LoginResponse>('/api/login', formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
    return resp.data;
}

export function logout() {
    removeToken();
    const navigate = useNavigate();
    navigate("/");
}

export async function getMe() {
    let resp = await axios.get<User>('/api/me');
    return resp.data;
}