import { useEffect, useContext } from 'react'

import { useNavigate, useParams } from "react-router-dom";

import API from 'Api'
import { useRequest } from 'Shared/Hooks';
import { User } from 'Shared/Models'
import SimplePage from 'Shared/SimplePage';
import { UserContainer } from 'Shared/UserContainer'

const CheckReset = () => {
  const { code } = useParams<{ code: string }>();
  const [loading, error, run] = useRequest<User>({} as User)
  const { userLoading, setUser } = useContext(UserContainer)
  const navigate = useNavigate();

  useEffect(() => {
    if (!userLoading) {
      // wait until defailt whoami returns before attempting reset
      run(API.checkReset(code as string), user => {
        if (user.id) {
          setUser(user);
          navigate("/change-password");
        }
      });
    }
  }, [run, userLoading, setUser, code, navigate])

  return (
    <SimplePage title='Logging you in...' loading={userLoading || loading} error={error} centered />
  )
}

export default CheckReset;
