import React, { useEffect, useState } from 'react';
import { getProfile } from '../services/api'; // Import the getProfile function

const Profile = () => {
  const [profile, setProfile] = useState(null);
  const [error, setError] = useState('');
  const token = localStorage.getItem('token'); // Get the JWT token from localStorage

  useEffect(() => {
    if (token) {
      // Fetch profile data only if token exists
      getProfile(token)
        .then((response) => {
          setProfile(response.data); // Assuming the API returns user profile data
        })
        .catch((err) => {
          setError('Error');
          console.error(err);
        });
    } else {
      setError('Login please');
    }
  }, [token]);

  if (error) {
    return <div>{error}</div>; // Display error if token is not available or there was an error fetching the profile
  }

  if (!profile) {
    return <div>Loading...</div>; // Show loading message while profile is being fetched
  }

  return (
    <div>
      <h2>Profile</h2>
      <p>Username: {profile.username}</p>
      <p>Email: {profile.email}</p>
    </div>
  );
};

export default Profile;
