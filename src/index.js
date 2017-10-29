import React from 'react';
import { BrowserRouter } from 'react-router-dom';
import { render } from 'react-dom';
import routes from './routes';
import registerServiceWorker from './registerServiceWorker';

render(<BrowserRouter>{routes}</BrowserRouter>, document.getElementById('app'));
registerServiceWorker();
