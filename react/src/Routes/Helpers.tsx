import { ReactNode, useContext } from "react";

import { Navigate } from "react-router-dom";
import { Loader, Container, Dimmer } from "semantic-ui-react";

import SimplePage from "Shared/SimplePage";
import { UserContainer } from "Shared/UserContainer";

type RequireAuthProps = {
  children: ReactNode;
  redirectTo: string;
};

// check the user is logged in, and redirect to login screen if still not auth'd
export function RequireAuth({ children, redirectTo }: RequireAuthProps) {
  const { user, userLoading } = useContext(UserContainer);

  if (userLoading) {
    return <BigLoader />;
  }

  if (!user.id) {
    return <Navigate to={redirectTo} replace />;
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
