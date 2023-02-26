import React, { ReactNode } from 'react'

import { Helmet } from 'react-helmet'
import { Container, Grid, Header, Icon, Message, SemanticICONS, Placeholder } from 'semantic-ui-react'

type SimplePageProps = React.PropsWithChildren<{
  title: string
  icon?: SemanticICONS
  centered?: boolean
  loading?: boolean
  error?: string
}>

const SimplePage = ({ title, icon, centered, loading, error, children }: SimplePageProps): JSX.Element => (
  <Container style={{ paddingTop: '7em' }}>
    {title && !loading &&
      <Helmet>
        <title>{title}</title>
      </Helmet>}

    <Content centered={!!centered}>
      <Header as='h1'>
        {icon && <Icon name={icon} />}{loading ? <PlaceholderTitle centered={centered}/> : title}
      </Header>

      {error && <Message negative>{error}</Message>}

      {loading ? <PlaceholderPost/> : children}
    </Content>
  </Container>
);

export default SimplePage;

const PlaceholderTitle = ({ centered }: { centered?: boolean}) => (
  <div style={{
    display:'table-cell',
    width: '10em',
    paddingTop: '0.4em', 
    paddingLeft: centered!! ? '' : '0.5em',
    }}>
    <Placeholder>
      <Placeholder.Header>
        <Placeholder.Line/>
      </Placeholder.Header>
    </Placeholder>
  </div>
)

const PlaceholderPost = () => (
  <Placeholder>
    <Placeholder.Paragraph>
      <Placeholder.Line/>
      <Placeholder.Line/>
      <Placeholder.Line/>
    </Placeholder.Paragraph>
  </Placeholder>
)

type ContentProps = {
  children: ReactNode;
  centered: boolean
}

const Content = ({ centered, children }: ContentProps): JSX.Element => {
  if (!centered) {
    return <>{children}</>;
  }

  return <>
    <Grid centered>
      <Grid.Column textAlign="center" mobile={16} tablet={8} computer={6}>
        {children}
      </Grid.Column>
    </Grid>
  </>
}
