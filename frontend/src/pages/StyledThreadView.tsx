import BasicCommentList from "../components/PostList";
import { Button, Card, CardContent, Fade, Typography } from "@mui/material";
import { Link } from "react-router-dom";
import Typewriter from "typewriter-effect";
import UserContext from "../contexts/UserContext";

import React, { useState } from "react";

const StyledThreadView: React.FC = () => {
    const [isShowTips, setIsShowTips] = useState(false);
    const { username } = React.useContext(UserContext);

    const showTips = () => {
        setIsShowTips(true);
    };

    return (
        <div style={{ width: "30vw", margin: "auto" }}>
            <Typography variant="h3" style={{ padding: "" }}>
                {"Welcome, " + username + "!"}
            </Typography>
            <Card>
                <CardContent>
                    <Typography component="p">{"Viewing thread:"}</Typography>
                    <Typography variant="h5" component="h5">
                        {"Cool thread"}
                    </Typography>
                    <Typography color="textSecondary" gutterBottom>
                        {"by Aiken"}
                    </Typography>
                    <Typography variant="body2" component="p">
                        {'"The best way to predict the future is to invent it."'}
                        <br />
                        {"- Alan Kay"}
                    </Typography>
                </CardContent>
            </Card>

            <BasicCommentList styled={true} />

            <Link to="/">
                <Button variant="contained" color="secondary">
                    {"Back to threads"}
                </Button>
            </Link>
        </div>
    );
};

export default StyledThreadView;
