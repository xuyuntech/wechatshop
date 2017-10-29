import { render } from 'react-dom';
import routes from './routes';
import registerServiceWorker from './registerServiceWorker';

render(routes, document.getElementById('app'));
registerServiceWorker();
