import axios from 'axios';

// Base URL for API
const API_URL = 'http://localhost:8080'; // Change this to your backend URL

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// User API calls
export const getUsers = () => api.get('/users');
export const getUserById = (id) => api.get(`/users/${id}`);
export const createUser = (user) => api.post('/users', user);

// Product API calls
export const getProducts = () => api.get('/products');
export const getProductById = (id) => api.get(`/products/${id}`);
export const createProduct = (product) => api.post('/products', product);

// Shopping Cart API calls
export const getShoppingCart = () => api.get('/shopping_cart');
export const createCartItem = (cartItem) => api.post('/shopping_cart/items', cartItem);

// Authentication API calls (Login and Register)
export const login = (username, password) => api.post('/login', { username, password });
export const register = (username, email, password) => api.post('/register', { username, email, password });

// Profile API call
export const getProfile = (token) => {
    return api.get('/profile', {
      headers: {
        Authorization: `Bearer ${token}`, // Send token as Authorization header
      },
    });
  };
