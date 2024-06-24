import ky from "ky";

export const apiClient = ky.create({ prefixUrl: process.env.API_BASE_URL });
