import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import NavBar from './components/NavBar'; // Import NavBar
import ProductList from './components/ProductList';
import ProductDetail from './components/ProductDetail';
import UserList from './components/UserList';
import UserDetail from './components/UserDetail';
import ShoppingCart from './components/ShoppingCart';
import Login from './components/Login'; // Import Login component
import Register from './components/Register'; // Import Register component
import Profile from './components/Profile'; // Import Profile component

const App = () => {
  const [token, setToken] = useState(localStorage.getItem('token') || '');

  // Update setToken to store token in localStorage
  const handleTokenChange = (newToken) => {
    setToken(newToken);
    localStorage.setItem('token', newToken);
  };

  return (
    <Router>
      <div>
        <NavBar token={token} /> {/* Pass token to NavBar */}

        <Routes>
          {/* Public Routes */}
          <Route exact path="/products" element={<ProductList />} />
          <Route exact path="/products/:id" element={<ProductDetail />} />
          <Route exact path="/users" element={<UserList />} />
          <Route exact path="/users/:id" element={<UserDetail />} />
          <Route exact path="/shopping_cart" element={<ShoppingCart />} />

          {/* Authentication Routes */}
          <Route exact path="/login" element={<Login setToken={handleTokenChange} />} />
          <Route exact path="/register" element={<Register />} />

          {/* Profile Route */}
          <Route exact path="/profile" element={<Profile />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
