import React, { useState, useEffect, useCallback } from 'react'

import API from 'Api'

import { User } from './Models';

export const UserContainer = React.createContext({
  user: {} as User,
  setUser: (user: User) => { },
  userLoading: false,
  setLoading: (loading: boolean) => { },
  handleLogout: () => { }
});

export const WithUser = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState({} as User)
  const [userLoading, setLoading] = useState(true)

  const handleLogout = useCallback(() => {
    API.logout()
      .then(() => {
        setUser({} as User);
      })
  }, [])

  useEffect(() => {
    API.whoami()
      .then(user => {
        setUser(user)
      })
      .finally(() => {
        setLoading(false)
      })
  }, [])

  return (
    <UserContainer.Provider value={{ user, setUser, userLoading, setLoading, handleLogout }}>
      {children}
    </UserContainer.Provider >
  );
}
