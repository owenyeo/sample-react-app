import React, {createContext, useState, ReactNode} from 'react';

interface UserContextType {
    username: string;
    setUsername: (username: string) => void;
    isLoggedIn: boolean;
    setIsLoggedIn: (isLoggedIn: boolean) => void;
    toggleIsLoggedIn: () => void;
};

const UserContext = createContext<UserContextType>({
    username: "",
    setUsername: () => {},
    isLoggedIn: false,
    setIsLoggedIn: () => {},
    toggleIsLoggedIn: () => {},
    }
);

type Props = {
    children: ReactNode;
};

export const UserProvider: React.FC<Props> = ({children}: Props) => {
    const [username, setUsername] = useState<string>("");
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    const toggleIsLoggedIn = () => {
        setIsLoggedIn(!isLoggedIn);
        setUsername("");
    }

    return (
        <UserContext.Provider value={{username, setUsername, isLoggedIn, setIsLoggedIn, toggleIsLoggedIn}}>
            {children}
        </UserContext.Provider>
    );
}

export const useUser = () => React.useContext(UserContext);

export default UserContext;
