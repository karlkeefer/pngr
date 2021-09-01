import React, { useState, useEffect, useCallback } from 'react'

import API from 'Api'

export const emptyUser = {
  id: 0,
  status: 0
};

export const User = React.createContext({
  user: emptyUser,
  setUser: () => {},
  loading: false,
  setLoading: ()=> {},
  handleLogout: () => {}
});


export const WithUser = ({children}) => {
  const [user, setUser] = useState(emptyUser)
  const [loading, setLoading] = useState(true)

  const handleLogout = useCallback(()=>{
    API.logout()
      .then(()=>{
        setUser(emptyUser);
      })
  }, [])

  useEffect(()=>{
    API.whoami()
      .then(user => {
        setUser(user)
      })
      .finally(() => {
        setLoading(false)
      })
  }, [setUser])
  
  return (
    <User.Provider value={{user, setUser, loading, setLoading, handleLogout}}>
      {children}
    </User.Provider >
  );
}
