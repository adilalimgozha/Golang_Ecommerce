import React, { useEffect, useState } from 'react';
import { getUsers } from '../services/api';
import { Link } from 'react-router-dom';

const UserList = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    getUsers()
      .then((response) => setUsers(response.data))
      .catch((error) => console.error('Error fetching users:', error));
  }, []);

  return (
    <div>
      <h1>User List</h1>
      <ul>
        {users.map((user) => (
          <li key={user.UserID}>
            <Link to={`/users/${user.UserID}`}>{user.Username}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default UserList;
