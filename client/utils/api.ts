import axios from 'axios';
import Cookies from 'js-cookie';

const API_BASE_URL = 'http://localhost:8080/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const login = async (email: string, password: string) => {
  const response = await api.post('/login', { email, password });

  if (response.data?.data.token) {
    Cookies.set('token', response.data.data.token);
  }
  return response.data;
};

export const register = async (firstName: string, lastName: string, email: string, password: string) => {
  const response = await api.post('/register', { first_name: firstName, last_name: lastName, email, password });
  return response.data;
};

export const verifyEmail = async (token: string) => {
  const response = await api.post('/verify-email', { token });
  if (response.data?.data.token) {
    Cookies.set('token', response.data.data.token);
  }
  return response.data;
};

export const validateToken = async () => {
  const token = Cookies.get('token');
  if (!token) {
    throw new Error('No token found');
  }
  try {
    const response = await api.get('/validate-token', {
      headers: { Authorization: `Bearer ${token}` },
    });

    return response.status === 200;
  } catch (error) {
    Cookies.remove('token');
    throw error;
  }
};

export const notifyForgotPassword = async (email: string) => {
  const response = await api.post('/notif-forgot-password', { email });
  return response.data;
};

export const resetPassword = async (token: string, password: string) => {
  const response = await api.post('/reset-password', { token, password });
  return response.data;
};

export const editUser = async (userData: {
  first_name: string;
  last_name: string;
  phone_number: string;
  address: string;
  address2: string;
  city: string;
  state: string;
  zipcode: string;
  profile_photo_url: string;
}) => {
  const token = Cookies.get('token');
  if (!token) {
    throw new Error('No token found');
  }

  const response = await api.put('/users/edit', userData, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return response.data;
};

export const getUserProfile = async () => {
  const token = Cookies.get('token');

  if (!token) {
    throw new Error('No token found');
  }

  const response = await api.get('/users/profile', {
    headers: { Authorization: `Bearer ${token}` },
  });

  if (response.status === 200) {
    return response.data.data;
  }

  throw new Error('Failed to fetch user profile');
};
