import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getUserById } from '../services/api';

const UserDetail = () => {
  const { id } = useParams();
  const [user, setUser] = useState(null);

  useEffect(() => {
    getUserById(id)
      .then((response) => setUser(response.data))
      .catch((error) => console.error('Error fetching user:', error));
  }, [id]);

  if (!user) return <div>Loading...</div>;

  return (
    <div>
      <h1>{user.Username}</h1>
      <p>Email: {user.Email}</p>
      {/* Add more user details here */}
    </div>
  );
};

export default UserDetail;
