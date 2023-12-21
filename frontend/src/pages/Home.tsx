import BasicThreadList from "../components/BasicThreadList";
import LogIn from "../components/LogIn";
import React from "react";

const Home: React.FC = () => {
    return (
        <>
            <h3>
                {"Welcomed to CVWO's sample react app! Here's a basic list of foum threads for you to experiment with."}
            </h3>
            <br />
            <LogIn />
        </>
    );
};

export default Home;
