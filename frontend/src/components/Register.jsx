import React, { useState } from 'react';
import { register } from '../services/api';  // Correctly import the register function

const Register = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  const handleRegister = async (e) => {
    e.preventDefault();

    try {
      // Using the imported register function
      const response = await register(username, email, password);
      setSuccessMessage(response.data.message); // Show success message on successful registration
      setError('');
    } catch (err) {
      setError(err.response?.data?.error || 'Ошибка при регистрации');
      setSuccessMessage('');
    }
  };

  return (
    <div>
      <h2>Registration</h2>
      <form onSubmit={handleRegister}>
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
          <label>Email:</label>
          <input
            type="email"
            placeholder="enter email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
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
        <button type="submit">Register</button>
      </form>

      {error && <p style={{ color: 'red' }}>{error}</p>} {/* Display error if any */}
      {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>} {/* Display success message */}
    </div>
  );
};

export default Register;
