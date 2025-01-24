import axios from "axios";
import useSWR from 'swr';
import { BaseResp } from '../model/resp';
import { ListUsersResponse, LoginRequest, LoginResponse, User } from "../model/user";
import { fetcher } from './common';
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
}

export async function getMe() {
    let resp = await axios.get<User>('/api/me');
    return resp.data;
}

export async function changePassword(newPassword: string) {
    const formData = new FormData();
    formData.append('password', newPassword);
    let resp = await axios.put<BaseResp>('/api/me/password', formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
    return resp.data;
}

export function useUsers(keyword: string, pageNo: number, pageSize: number, shouldFetch = true) {
    const key = shouldFetch ? `/api/users?keyword=${keyword}&page_no=${pageNo}&page_size=${pageSize}` : null;
    return useSWR<ListUsersResponse>(key, fetcher);
}

export async function listUsers(keyword: string, pageNo: number, pageSize: number) {
    const resp = await axios.get<ListUsersResponse>(`/api/users?keyword=${keyword}&page_no=${pageNo}&page_size=${pageSize}`);
    return resp.data;
}
