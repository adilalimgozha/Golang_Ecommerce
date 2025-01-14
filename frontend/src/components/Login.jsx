import React, { useState } from 'react';
import { login } from '../services/api';  // Import the login function from api.js

const Login = ({ setToken }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await login(username, password); // Call the login function from api.js
      const token = response.data.token;

      // Store the token in localStorage and update the state with the new token
      setToken(token);
      localStorage.setItem('token', token); // Store token in localStorage
      setError('');
    } catch (err) {
      setError('Неверные учетные данные');  // Show error if login fails
    }
  };

  return (
    <div>
      <h2>Login</h2>
      <form onSubmit={handleLogin}>
        <div>
          <label>Username:</label>
          <input
            type="text"
            placeholder="enter username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            placeholder="enter password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit">Login</button>
      </form>

      {error && <p style={{ color: 'red' }}>{error}</p>}  {/* Show error message if login fails */}
    </div>
  );
};

export default Login;
