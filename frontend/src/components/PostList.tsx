import React, { useState, useEffect } from "react";
import axios from "axios";
import PostItem from "./PostItem";
import { Button, TextField } from "@mui/material";
import Post from "../types/Post";
import { useUser } from "../contexts/UserContext";

type Props = {
    styled: boolean;
};

const BasicPostsList: React.FC<Props> = ({ styled }: Props) => {
    const [posts, setPosts] = useState<Post[]>([]);
    const [newContent, setNewContent] = useState<string>("");
    const [newTitle, setNewTitle] = useState<string>("");
    const { username } = useUser();

    const fetchPosts = () => {
        axios.get("http://localhost:8000/posts")
        .then(response => {
            setPosts(response.data.payload.data || []);
        })
        .catch(error => {
            console.error('There was an error fetching the posts!', error);
        });
    };
    
    useEffect(() => {
        fetchPosts();
    }, []);  // Dependency array is empty to run only on mount

    const handleNewPostChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewContent(event.target.value);
    };

    const handleNewPostSubmit = (event: React.FormEvent) => {
        event.preventDefault();
        axios.post('http://localhost:8000/posts/new', {
            id: posts.length + 1,
            title: newTitle,
            content: newContent,
            author: username, 
            timestamp: new Date(),
        })
        .then(response => {
            console.log("New post response data:", response.data);
            setPosts(prevPosts => [...prevPosts, response.data]);
            setNewContent('');
            setNewTitle('');
            fetchPosts();
        })
        .catch(error => {
            console.error('Error adding post', error);
        });
    };

    const handleNewTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewTitle(event.target.value);
    };

    return (
        <div>
            {posts.map((post) => (
                <PostItem styled={styled} Post={post} />
            ))}
            <form onSubmit={handleNewPostSubmit} style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
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
                <Button variant="contained" type="submit" style={{ marginTop: '10px' }}>Submit</Button>
            </form>
        </div>
    );
};

export default BasicPostsList;
