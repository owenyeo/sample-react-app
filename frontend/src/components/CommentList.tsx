import CommentItem from "./CommentItem";
import Comment from "../types/Comment";
import React, { useState } from "react";
import { Button, TextField } from "@mui/material"; // Import Button from your UI library

type Props = {
    styled: boolean;
};

const BasicCommentList: React.FC<Props> = ({ styled }: Props) => {
    const [comments, setComments] = useState<Comment[]>([
        {
            body:
                "Any fool can write code that a computer can understand.\n" +
                "Good programmers write code that humans can understand.\n" +
                " ~ Martin Fowler",
            author: "Benedict",
            timestamp: new Date(2022, 10, 28, 10, 33, 30),
        },
        {
            body: "Code reuse is the Holy Grail of Software Engineering.\n" + " ~ Douglas Crockford",
            author: "Casey",
            timestamp: new Date(2022, 11, 1, 11, 11, 11),
        },
        {
            body: "Nine people can't make a baby in a month.\n" + " ~ Fred Brooks",
            author: "Duuet",
            timestamp: new Date(2022, 11, 2, 10, 30, 0),
        },
    ]);

    const [newComment, setNewComment] = useState<string>("");

    const handleNewCommentChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewComment(event.target.value);
    }

    const addComment = () => {
        if (newComment.trim() !== "") {
            const comment: Comment = {
                body: newComment,
                author: "User",
                timestamp: new Date(),
            };
            setComments([...comments, comment]);
            setNewComment("");
        }
    };

    return (
        <div>
            <ul>
                {comments.map((comment, index) => (
                    <CommentItem comment={comment} styled={styled} key={index.toString()} />
                ))}
            </ul>
            <TextField
                label="New Post"
                variant="outlined"
                onChange={handleNewCommentChange}
                value={newComment}
            />
            <Button variant="contained" color="primary" onClick={addComment}>
                Add Comment
            </Button>
        </div>
    );
};

export default BasicCommentList;
