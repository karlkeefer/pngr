import { ReactNode, useContext } from "react";

import { Navigate, useLocation } from "react-router-dom";
import { Loader, Container, Dimmer } from "semantic-ui-react";

import SimplePage from "Shared/SimplePage";
import { UserContainer } from "Shared/UserContainer";

type RequireAuthProps = {
  children: ReactNode;
};

// check the user is logged in, and redirect to login screen if still not auth'd
export const RequireAuth = ({ children }: RequireAuthProps) => {
  const { user, userLoading } = useContext(UserContainer);
  const { pathname } = useLocation();

  if (userLoading) {
    return <BigLoader />;
  }

  if (!user.id) {
    return <Navigate to='/login' state={{from: pathname}} replace />;
  }

  return <>{children}</>
}

export const NoMatch = () => (
  <SimplePage icon='cancel' title='Not Found'>
    <p>The page you are trying to view does not exist!</p>
  </SimplePage>
);

const BigLoader = () => (
  <Container>
    <Dimmer active inverted>
      <Loader size="big" />
    </Dimmer>
  </Container>
)
