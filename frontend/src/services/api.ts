import axios, { AxiosInstance } from 'axios';

const API_BASE_URL = 'http://localhost:8000';

const api: AxiosInstance = axios.create({
    baseURL: API_BASE_URL,
    });

export const addUser = async (userData: any) => {
    try {
        const response = await api.post('/addUser', userData);
        return response.data; // Return the response data if needed
    } catch (error) {
        throw error; // Throw an error to handle it in the component
    }
};

export const getUserById = async (userId: number) => {
    try {
        const response = await api.get(`/getUser/${userId}`);
        return response.data; // Return the response data if needed
    } catch (error) {
        throw error; // Throw an error to handle it in the component
    }
};