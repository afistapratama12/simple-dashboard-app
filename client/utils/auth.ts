import { validateToken } from './api';

export const isAuthenticated = async () => {
  try {
    await validateToken();
    return true;
  } catch (error) {
    return false;
  }
};
