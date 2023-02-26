import { useEffect, useContext } from 'react'

import { useNavigate, useParams } from "react-router-dom";
import { Message } from 'semantic-ui-react'

import API from 'Api'
import { useRequest } from 'Shared/Hooks';
import { User } from 'Shared/Models'
import SimplePage from 'Shared/SimplePage'
import { UserContainer } from 'Shared/UserContainer'

const Verify = () => {
  const { code } = useParams<{ code: string }>();
  const [loading, error, run, user] = useRequest<User>({} as User)
  const { userLoading, setUser } = useContext(UserContainer)
  const navigate = useNavigate();

  useEffect(() => {
    if (!userLoading) {
      // wait until default whoami (called within the UserContainer) returns 
      // before attempting reset, otherwise there is a race condition
      if (!code) {
        navigate("/")
        return
      }

      run(API.verify({ code }), user => {
        setUser(user);
        setTimeout(() => {
          navigate("/posts")
        }, 2500);
      })
    }
  }, [run, setUser, code, userLoading, navigate])

  return (
    <SimplePage title='Account Verification' centered loading={loading} error={error}>
      {user && user.id && user.id > 0 && 
        <Message positive>Success! You have verified your email!</Message>}
    </SimplePage>
  );
}

export default Verify;
