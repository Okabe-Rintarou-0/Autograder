import { BaseResp, ListItemResponse } from "./resp";

export const CommonUser = 1;
export const Administrator = 2;

export interface LoginRequest {
    identifier: string;
    password: string;
}

export interface LoginResponse extends BaseResp {
    token: string;
}

export interface User extends BaseResp {
    user_id: number;
    real_name: string;
    role: number;
    username: string;
    email: string;
}

export interface ListUsersResponse extends ListItemResponse<User> { }
