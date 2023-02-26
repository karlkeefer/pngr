import { useEffect, useCallback, useContext } from 'react'

import { useLocation, useNavigate } from "react-router-dom";
import { Form, Button, Message } from 'semantic-ui-react'

import API from 'Api'
import { useRequest, useFields } from 'Shared/Hooks';
import { User } from 'Shared/Models'
import { UserContainer } from 'Shared/UserContainer'

const LogInForm = () => {
  const location = useLocation()
  const { user, setUser } = useContext(UserContainer)
  const [loading, error, run] = useRequest({} as User)
  const { fields, handleChange, setFields } = useFields({} as User)
  const navigate = useNavigate();

  useEffect(() => {
    if (user.id) {
      navigate("/posts", { replace: true})
    }
  }, [user, navigate])

  const handleSubmit = useCallback(() => {
    run(API.login(fields), user => {
      if (user.id) {
        setUser(user);
        setFields({} as User)
        const { from } = location.state || { from: { pathname: "/posts" } };
        navigate(from);
      }
    });
  }, [run, fields, setFields, setUser, navigate, location.state])

  return (
    <Form error name="login" loading={loading} onSubmit={handleSubmit}>
      <Message error>{error}</Message>
      <Form.Input
        autoFocus
        size="big"
        name="email"
        type="email"
        placeholder="Email"
        required
        onChange={handleChange} />
      <Form.Input
        size="big"
        name="pass"
        type="password"
        placeholder="Password"
        required
        onChange={handleChange} />
      <Button primary fluid size="huge" type="submit">Log In</Button>
    </Form>
  )
}

export default LogInForm;
