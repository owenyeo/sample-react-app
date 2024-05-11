import React from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import { Link } from 'react-router-dom';
import { useUser } from '../contexts/UserContext';
import { toggleButtonGroupClasses } from '@mui/material';

const NavBar: React.FC = () => {
    const { username, toggleIsLoggedIn } = useUser();
    return (
        <AppBar position="static">
            <Toolbar>
                <Typography variant="h5" style={{ flexGrow: 1 }}>
                    {"CVWO Forum"}
                </Typography>
                <Button color="inherit" component={Link} to="/">
                    {"Home"}
                </Button>
                <Button color="inherit" component={Link} to="/thread/1/styled">
                    {"Thread"}
                </Button>
                {username && (
                    <Button color="inherit" onClick={toggleIsLoggedIn}>
                        Logout
                    </Button>
                )}
            </Toolbar>
        </AppBar>
    );
}

export default NavBar;