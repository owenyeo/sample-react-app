import PostItem from "./PostItem";
import axios from "axios";
import Post from "../types/Post";
import React, { useState, useEffect } from "react";
import { Button, TextField } from "@mui/material"; // Import Button from your UI library

type Props = {
    styled: boolean;
};

const BasicPostsList: React.FC<Props> = ({ styled }: Props) => {
    const [Posts, setPost] = useState<Post[]>([]);
    const [newContent, setNewContent] = useState<string>("");
    const [newTitle, setNewTitle] = useState<string>("");


    useEffect(() => {
        axios.get("http://localhost:8000/posts")
        .then((response) => {
            console.log(response.data);
            if (response.data.payload.data) {
            setPost(response.data.payload.data);
            } else {
                setPost([]);
            }
        })
        .catch((error) => {
            console.error('There was an error!', error);
        });
    }, []);

    const handleNewPostChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewContent(event.target.value);
    }

    const handleNewPostSubmit = (event: React.FormEvent) => {        
        axios.post('http://localhost:8000/posts/new', {
            id: Posts.length + 1,
            title: newTitle,
            content: newContent,
            author: 'User', // Replace with the actual author name
            timestamp: new Date(),
        })
        .then(response => {
            // Add the new comment to the state
            setPost(prevPosts=> [...prevPosts, response.data]);
            setNewContent('');
        })
        .catch(error => {
            console.error('Error adding comment', error);
        });
    };

    const handleNewTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewTitle(event.target.value);
    }

    return (
        <div>
            {Posts && Posts.map((post, index) => (
                <PostItem key={index} styled={styled} Post={post} />
            ))}
            <form onSubmit={handleNewPostSubmit}>
                <TextField
                    value={newTitle}
                    onChange={handleNewTitleChange}
                    placeholder="Write a new title..."
                />
                <TextField
                    value={newContent}
                    onChange={handleNewPostChange}
                    placeholder="Write a new post..."
                />
                <Button type="submit"  style={{ display: 'block', marginTop: '10px' }}>Submit</Button>
            </form>
        </div>
    );
};

export default BasicPostsList;
