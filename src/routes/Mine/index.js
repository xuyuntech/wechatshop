import React from 'react';
import Layout from 'components/Layout';
import Loadable from 'util/Loadable';

const LoadableComponent = Loadable(
  import(/* webpackChunkName: 'mine' */ './Mine'),
);

export default class extends React.Component {
  render() {
    return (
      <Layout>
        <LoadableComponent />
      </Layout>
    );
  }
}
