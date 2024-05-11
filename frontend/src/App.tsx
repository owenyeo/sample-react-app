import Home from "./pages/Home";
import StyledThreadView from "./pages/StyledThreadView";
import React from "react";
import "./App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { blue, orange } from "@mui/material/colors";
import { UserProvider } from "./contexts/UserContext";
import NavBar from "./components/NavBar";

const theme = createTheme({
    palette: {
        primary: blue,
        secondary: orange,
    },
});

const App: React.FC = () => {
    return (
        <UserProvider>
            <div className="App">
                <ThemeProvider theme={theme}>
                    <BrowserRouter>
                        <NavBar/>
                        <Routes>
                            <Route path="/thread/1" element={<StyledThreadView />} />
                            <Route path="/thread/1/styled" element={<StyledThreadView />} />
                            <Route path="/" element={<Home />} />
                        </Routes>
                    </BrowserRouter>
                    
                </ThemeProvider>
            </div>
        </UserProvider>
    );
};

export default App;
