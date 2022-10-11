import React, { useContext } from 'react'

import { UserContainer } from 'Shared/UserContainer'

export const LoggedIn = ({ children }: { children: React.ReactNode }) => {
  const { user } = useContext(UserContainer)
  return <>{user.id && user.id > 0 && children}</>
}

export const Anon = ({ children }: { children: React.ReactNode }) => {
  const { user } = useContext(UserContainer)
  return <>{(user.id === 0 || !user.id) && children}</>
}
