import React from 'react';
import ReactDOM from 'react-dom';

import 'semantic-ui-css/semantic.min.css';
import './index.css';

import App from './App';
// import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<App />, document.getElementById('root'));
// if you want to run this as a progressive webapp, you can register this service worker, 
// but you also need to add handling for upgrades otherwise folks will get stuck with old versions of your app indefinitely
// registerServiceWorker();

// unregister all serviceWorkers
if ('serviceWorker' in navigator) {
  navigator.serviceWorker.getRegistrations().then(function(registrations) {
    for(let registration of registrations) {
      registration.unregister();
    }
  });
}
