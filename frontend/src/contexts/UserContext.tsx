import React, {createContext, useState, ReactNode} from 'react';

interface UserContextType {
    username: string;
    setUsername: (username: string) => void;
};

const UserContext = createContext<UserContextType>({
    username: "",
    setUsername: () => {},
});

type Props = {
    children: ReactNode;
};

export const UserProvider: React.FC<Props> = ({children}: Props) => {
    const [username, setUsername] = useState<string>("");

    return (
        <UserContext.Provider value={{username, setUsername}}>
            {children}
        </UserContext.Provider>
    );
}

export default UserContext;
