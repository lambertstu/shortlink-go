import axios from "axios";
import qs from "qs";

const API_BASE_URL = "http://localhost:8001/v1/group"; // 你的后端地址

// 创建群组
export function createGroup(data: { name: string; username: string }) {
  return axios.post(`${API_BASE_URL}/create`, qs.stringify(data), {
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
  });
}

// 更新群组
export function updateGroup(data: { gid: string; name: string; username: string }) {
  return axios.post(`${API_BASE_URL}/update`, qs.stringify(data), {
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
  });
}

// 删除群组
export function deleteGroup(data: { gid: string; username: string }) {
  return axios.post(`${API_BASE_URL}/delete`, qs.stringify(data), {
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
  });
}

// 获取群组信息
export function getGroupInfo(username: string) {
  return axios.get(`${API_BASE_URL}/info`, {
    params: { username },
  });
}
