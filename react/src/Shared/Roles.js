import { useContext } from 'react'

import { User } from 'Shared/Context'

export const LoggedIn = ({children}) => {
  const {user} = useContext(User)
  return user.id > 0 ? children : false;
}

export const Anon = ({children}) => {
  const {user} = useContext(User)
  return user.id === 0 ? children : false;
}
