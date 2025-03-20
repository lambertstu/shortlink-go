import axios from "axios";
import qs from "qs";

const API_BASE_URL = "http://localhost:8888/v1/user"; // 你的后端地址

// 用户注册
export function registerUser(data: {
  email: string;
  password: string;
  phone: string;
  realName: string;
  username: string;
}) {
  return axios.post(`${API_BASE_URL}/register`, data, {
    headers: { "Content-Type": "application/json" },
  });
}

// 用户登录
export function loginUser(data: { username: string; password: string }) {
  return axios.post(`${API_BASE_URL}/login`, qs.stringify(data), {
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
  });
}

// 用户登出
export function logoutUser(data: { token: string; username: string }) {
  return axios.post(`${API_BASE_URL}/logout`, qs.stringify(data), {
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
  });
}

// 查询用户是否存在
export function checkUserExist(username: string) {
  return axios.get(`${API_BASE_URL}/exist`, { params: { username } });
}

// 查询用户登录状态
export function checkLoginStatus(data: { token: string; username: string }) {
  return axios.get(`${API_BASE_URL}/login-status`, { params: data });
}

// 更新或插入用户信息
export function upsertUser(data: {
  username: string;
  password: string;
  realName: string;
  phone: string;
  email: string;
}) {
  return axios.post(`${API_BASE_URL}/upsert`, qs.stringify(data), {
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
  });
}

// 删除用户
export function deleteUser(data: { username: string }) {
  return axios.post(`${API_BASE_URL}/delete`, qs.stringify(data), {
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
  });
}

// 获取用户信息
export function getUserInfo(username: string) {
  return axios.get(`${API_BASE_URL}/info`, { params: { username } });
}
