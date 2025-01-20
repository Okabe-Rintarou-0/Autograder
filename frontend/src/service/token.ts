import axios from 'axios';

export function loadToken() {
    const token = localStorage.getItem("token");
    if (token) {
        axios.defaults.headers.token = token;
    }
    console.log("Load token", token);
}

export function removeToken() {
    localStorage.removeItem("token");
    axios.defaults.headers.token = null;
}

export function setToken(token: string) {
    localStorage.setItem("token", token)
    axios.defaults.headers.token = token
}