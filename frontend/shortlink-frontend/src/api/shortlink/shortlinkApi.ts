import axios from "axios";

const API_BASE_URL = "http://localhost:8888/v1/shortlink"; // 你的后端地址

// 查询短链接列表
export function fetchShortLinks(params: { gid: string; page: number; size: number; orderTag: number }) {
  return axios.get(`${API_BASE_URL}/page`, { params });
}

// 创建短链接
export function createShortLink(data: { origin_url: string; gid: string; describe: string }) {
  return axios.post(`${API_BASE_URL}/create`, data);
}

// 更新短链接
export function updateShortLink(data: {
  gid: string;
  origin_url: string;
  short_uri: string;
  describe: string;
  favicon: string;
  clickNum: number;
  totalPv: number;
  totalUv: number;
  totalUip: number;
  todayPv: number;
  todayUv: number;
  todayUip: number;
}) {
  return axios.post(`${API_BASE_URL}/update`, data);
}

// 删除短链接
export function deleteShortLink(data: { shortUri: string; origin_url: string }) {
  return axios.post(`${API_BASE_URL}/delete`, data);
}
