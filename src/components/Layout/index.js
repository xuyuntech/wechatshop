import React from 'react';
import PropTypes from 'prop-types';
import { TabBar, NavBar } from 'xyui';
import 'font-awesome/css/font-awesome.css';
import cls from './styles';
import './layout.css';

/* eslint-disable no-console */

export default class extends React.Component {
  static propTypes = {
    children: PropTypes.any, // eslint-disable-line
  };
  static defaultProps = {
    children: null,
  };
  static contextTypes = {
    router: PropTypes.shape({
      history: PropTypes.shape({
        push: PropTypes.func.isRequired,
        replace: PropTypes.func.isRequired,
        createHref: PropTypes.func.isRequired,
      }).isRequired,
    }).isRequired,
  };
  constructor(props) {
    super(props);
    this.state = {
      type: 'home',
    };
  }
  switchTab = (to, type) => {
    this.setState({
      type,
    });
    const { history } = this.context.router;
    history.push(to);
  };
  render() {
    const path = this.context.router.route.match.path;
    return (
      <div>
        <div style={{ ...cls.navBar }}>
          <NavBar
            mode="light"
            icon={<i className="fa fa-angle-left" />}
            onLeftClick={() => console.log('onLeftClick')}
          >
            NavBar
          </NavBar>
        </div>
        <div>{this.props.children}</div>
        <div style={{ ...cls.tabBar }}>
          <TabBar
            unselectedTintColor="#949494"
            tintColor="#33A3F4"
            barTintColor="white"
          >
            <TabBar.Item
              title="首页"
              key="home"
              selected={path === '/'}
              icon={<i className="fa fa-home" />}
              selectedIcon={
                <i style={{ ...cls.iconS }} className="fa fa-home" />
              }
              onPress={() => {
                this.switchTab('/', 'home');
              }}
            />
            <TabBar.Item
              title="分类"
              key="category"
              selected={path === '/category'}
              icon={<i className="fa fa-home" />}
              selectedIcon={
                <i style={{ ...cls.iconS }} className="fa fa-home" />
              }
              onPress={() => {
                this.switchTab('/category', 'category');
              }}
            />
            <TabBar.Item
              title="我的"
              key="mine"
              selected={path === '/mine'}
              icon={<i className="fa fa-home" />}
              selectedIcon={
                <i style={{ ...cls.iconS }} className="fa fa-home" />
              }
              onPress={() => {
                this.switchTab('/mine', 'mine');
              }}
            />
            <TabBar.Item
              title="Demo"
              key="demo"
              selected={path === '/demo'}
              icon={<i className="fa fa-home" />}
              selectedIcon={
                <i style={{ ...cls.iconS }} className="fa fa-home" />
              }
              onPress={() => {
                this.switchTab('/demo', 'demo');
              }}
            />
          </TabBar>
        </div>
      </div>
    );
  }
}
