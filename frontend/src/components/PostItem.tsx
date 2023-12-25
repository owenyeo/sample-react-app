import Post from "../types/Post";

import React from "react";
import { Card, CardContent, Typography } from "@mui/material";
import { makeStyles } from "@mui/styles";

type Props = {
    Post: Post;
    styled: boolean;
};
const useStyles = makeStyles(() => ({
    PostBody: {
        fontSize: 16,
        whiteSpace: "pre-wrap",
        paddingBottom: "1em",
    },
    PostCard: {
        marginBottom: "1em",
        width: "100%",
        maxWidth: "1600px",
    },
    metadata: {
        fontSize: 16,
    },
}));

const PostItem: React.FC<Props> = ({ Post, styled }) => {
    const classes = useStyles();

    const formatDate = (date : Date) => {
        if (!(date instanceof Date)) {
            date = new Date(date); // Try to create a Date object from the provided value
        }
    
        if (isNaN(date.getTime())) {
            return 'Invalid Date'; // Handle cases where the date is invalid
        }

        const options = {
            year: 'numeric' as const,
            month: 'short' as const,
            day: 'numeric' as const,
            hour: 'numeric' as const,
            minute: 'numeric' as const,
            second: 'numeric' as const,
            hour12: true,
          };
        
          return new Intl.DateTimeFormat('en-US', options).format(date);
    }

    if (styled) {
        return (
            <Card className={classes.PostCard}>
                <CardContent>
                    <Typography variant="h5" component="h2">
                        {Post.title}
                    </Typography>
                    <Typography variant="body2" color="textPrimary" className={classes.PostBody} component="p">
                        {Post.content}
                    </Typography>
                    <Typography color="textSecondary" className={classes.metadata} gutterBottom>
                        {"Posted by " + Post.author + " on " + formatDate(Post.date)}
                    </Typography>
                </CardContent>
            </Card>
        );
    }

    // unstyled
    return (
        <li className={classes.PostBody}>
            {Post.content}
            <br />
            <em>{"posted by " + Post.author + " on " + Post.date.toLocaleString()}</em>
        </li>
    );
};

export default PostItem;
