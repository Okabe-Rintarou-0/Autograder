import { BaseResp, ListItemResponse } from "./resp";

export const CommonUser = 1;
export const Administrator = 2;

export const registerUserURL = "/api/register";
export const importCanvasUsersURL = "/api/register/canvas";

export interface LoginRequest {
    identifier: string;
    password: string;
}

export interface LoginResponse extends BaseResp {
    token: string;
}

export interface User extends BaseResp {
    id: number;
    real_name: string;
    role: number;
    username: string;
    email: string;
}

export interface ListUsersResponse extends ListItemResponse<User> { }

export interface RegisterUserRequest {
    username: string;
    real_name: string;
    email: string;
    password: string;
}

export interface ImportCanvasUsersRequest {
    course_id: number;
}