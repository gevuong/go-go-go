import ReactDOM from 'react-dom';
import React from 'react';
import App from './frontend/app';

document.addEventListener("DOMContentLoaded", () => {
    var root = document.getElementById("root")
    ReactDOM.render(<App />, root)
})
