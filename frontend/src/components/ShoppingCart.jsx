import React, { useEffect, useState } from 'react';
import { getShoppingCart } from '../services/api'; // Assuming getShoppingCart is imported from api.js

const ShoppingCart = () => {
  const [cartItems, setCartItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Fetch shopping cart data when the component mounts
  useEffect(() => {
    getShoppingCart()
      .then((response) => {
        setCartItems(response.data); // Set the cart items with product details
        setLoading(false);
      })
      .catch((err) => {
        setError('Failed to fetch shopping cart');
        setLoading(false);
      });
  }, []); // Empty dependency array ensures this runs once when the component mounts

  if (loading) {
    return <div>Loading your cart...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div>
      <h1>Your Shopping Cart</h1>
      {cartItems.length === 0 ? (
        <p>Your shopping cart is empty.</p>
      ) : (
        <div>
          <ul>
            {cartItems.map((item) => (
              <li key={item.CartItemID}>
                {/* Ensure Product is available */}
                {item.Product ? (
                  <div>
                    <strong>{item.Product.Name}</strong>
                    <div>Quantity: {item.Quantity}</div>
                    <div>Price: ${item.Product.Price}</div>
                    <div>Total: ${(item.Quantity * item.Product.Price).toFixed(2)}</div>
                  </div>
                ) : (
                  <div>
                    <strong>Product details are missing</strong>
                  </div>
                )}
              </li>
            ))}
          </ul>
          <div>
            <h3>
              Total Price: $
              {cartItems
                .reduce((acc, item) => acc + (item.Quantity * (item.Product?.Price || 0)), 0)
                .toFixed(2)}
            </h3>
          </div>
        </div>
      )}
    </div>
  );
};

export default ShoppingCart;
