import axios from "axios";

const backendUrl = import.meta.env.VITE_BACKEND_URL ?? "http://localhost:8080";

export const apiInstance = axios.create({
  baseURL: `${backendUrl}/api`,
});
