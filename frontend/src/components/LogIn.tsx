import "../App.css";

import React from "react";
import { Link, useNavigate} from "react-router-dom";
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import UserContext from "../contexts/UserContext";


const LogIn: React.FC = () => {
    const [name, setName] = React.useState('');
    const navigate = useNavigate();
    const { username, setUsername } = React.useContext(UserContext);

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setName(event.target.value);
    }

    const handleLogin = () => {
      // Perform login action here using 'name' state
      // For instance, make an API call to your backend with the username

      // Example API call using fetch
      fetch('http://localhost:8000/users/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ name }),
      })
      .then(response => {
          if (response.ok) {
              // Login successful
              console.log('Login successful!');
              setUsername(name);
              navigate('/thread/1/styled');
          } else {
              console.error('User not found!');
              handleRegistration();
          }
      })
      .catch(error => {
          console.error('Error during login:', error);
      });
    };

    const handleRegistration = () => {
      fetch('http://localhost:8000/users/new', {
                  method: 'POST',
                  headers: {
                    'Content-Type': 'application/json',
                  },
                  body: JSON.stringify({ name }),
              })
              .then(response => {
                  if (response.ok) {
                      // Registration successful
                      console.log('Registration successful!');
                      setUsername(name);
                      navigate('/thread/1/styled');
                  } else {
                      console.error('Registration failed!');
                  }
              })
              .catch(error => {
                  console.error('Error during Registration:', error);
              });
            };
    
    return (
        <Box
        component="form"
        sx={{
          '& .MuiTextField-root': { m: 1, width: '25ch' },
        }}
        noValidate
        autoComplete="off"
      >
        <div>
          <TextField
            required
            id="outlined-required"
            label="name"
            onChange={handleInputChange}
          />
        </div>
        <div>
            <Button onClick={handleLogin}>
                {"Log In"}
            </Button>
        </div>
      </Box>
    );
};

export default LogIn;