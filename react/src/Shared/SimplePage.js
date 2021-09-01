import React from 'react'
import { Container, Grid, Header, Icon, Message, Loader } from 'semantic-ui-react'
import { Helmet } from 'react-helmet'

const SimplePage = ({title, icon, centered, loading, error, children}) => (
  <Container style={{paddingTop: '7em'}}>
    {title && !loading ? 
      <Helmet>
        <title>{title}</title>
      </Helmet>
      : false}

    <Content centered={centered}>
      <Header as='h1'>
        {icon ? <Icon name={icon}/> : false }{title}
      </Header>

      {error ? <Message negative>{error}</Message> : false}

      {loading ? <Loader active inline='centered'/> : children}
    </Content>
  </Container>
);

export default SimplePage;

const Content = ({centered, children}) => (
  centered ? 
  <Grid centered>
    <Grid.Column textAlign="center" mobile={16} tablet={8} computer={6}>
      {children}
    </Grid.Column>
  </Grid>
  : children
)
