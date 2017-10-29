import React from 'react';
import { BrowserRouter, Route } from 'react-router-dom';
import Home from './Home';
import Category from './Category';
import Mine from './Mine';
import Demo from './Demo';

export default (
  <BrowserRouter>
    <div>
      <Route exact path="/" component={Home} />
      <Route path="/category" component={Category} />
      <Route path="/mine" component={Mine} />
      <Route path="/demo" component={Demo} />
    </div>
  </BrowserRouter>
);
