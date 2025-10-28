import { SERVER_BASE_URL } from "../config/config";

/*export const checkToken = async () => {
  try {
    const response = await fetch(`${SERVER_BASE_URL}/token`, {
      method: 'GET',
      credentials: 'include', // Для работы с HttpOnly cookies
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return {
      isAuthenticated: response.status === 200,
      status: response.status
    };
  } catch (error) {
    return {
      isAuthenticated: false,
      error: error.message
    };
  }
}; */

export const checkToken = () => {
    return {
      isAuthenticated: false,
      status: 200
    };
}