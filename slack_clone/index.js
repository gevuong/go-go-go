import ReactDOM from 'react-dom';
import React from 'react';
import App from './frontend/app.js';

document.addEventListener("DOMContentLoaded", () => {
    var root = document.getElementById("root")
    ReactDOM.render(<App />, root)
})
